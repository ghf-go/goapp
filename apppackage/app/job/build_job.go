package job

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ghf-go/goapp/apppackage/app/models"
	"github.com/ghf-go/goapp/apppackage/app/utils"
	"gopkg.in/yaml.v3"
)

// DepConfig 对应被构建仓库 .dep/dep.yaml 的结构
type DepConfig struct {
	Steps []struct {
		Name string `yaml:"name"`
		Cmd  string `yaml:"cmd"`
	} `yaml:"steps"`
	Artifacts []struct {
		Path string `yaml:"path"`
		Key  string `yaml:"key"`
	} `yaml:"artifacts"`
	Deploy struct {
		Name      string `yaml:"name"`
		Container string `yaml:"container"`
		Image     string `yaml:"image"`
	} `yaml:"deploy"`
}

// RunBuild 异步执行一次完整构建流程
func RunBuild(recordID uint64) {
	var record models.BuildRecordModel
	if err := utils.DB.First(&record, recordID).Error; err != nil {
		utils.Error("构建记录不存在 id=%d: %v", recordID, err)
		return
	}
	var project models.ProjectModel
	if err := utils.DB.First(&project, record.ProjectID).Error; err != nil {
		failBuild(&record, "工程不存在: %v", err)
		return
	}
	setStatus(&record, models.BuildStatusRunning)
	appendLog(&record, "开始构建 工程=%s 环境=%s 分支=%s", project.Name, record.Env, record.Branch)

	workDir := filepath.Join(utils.Conf.Build.WorkDir, strconv.FormatUint(record.ID, 10))
	defer os.RemoveAll(workDir)

	dep, ok := prepareCode(&record, &project, workDir)
	if !ok {
		return
	}
	if !runSteps(&record, workDir, dep) {
		return
	}
	version := bumpVersion(currentVersion(&project, record.Env))
	appendLog(&record, "新版本号: %s", version)

	if !uploadArtifacts(&record, &project, workDir, dep, version) {
		return
	}
	if !deployIfNeeded(&record, &project, dep, version) {
		return
	}
	saveProjectVersion(&project, record.Env, version)
	record.Version = version
	setStatus(&record, models.BuildStatusSuccess)
	if err := utils.DB.Model(&record).Update("version", version).Error; err != nil {
		utils.Error("回写构建版本号失败 record=%d: %v", record.ID, err)
	}
	appendLog(&record, "构建成功 版本=%s", version)
}

// prepareCode 拉取代码并解析 .dep/dep.yaml
func prepareCode(record *models.BuildRecordModel, project *models.ProjectModel, workDir string) (*DepConfig, bool) {
	appendLog(record, "拉取代码 %s (分支 %s)", project.GitUrl, record.Branch)
	cmd := fmt.Sprintf("git clone --depth 1 -b %s %s %s", record.Branch, project.GitUrl, workDir)
	if out, err := utils.RunShell("", buildTimeout(), cmd); err != nil {
		failBuild(record, "拉取代码失败: %v\n%s", err, out)
		return nil, false
	}
	data, err := os.ReadFile(filepath.Join(workDir, ".dep", "dep.yaml"))
	if err != nil {
		failBuild(record, "读取 .dep/dep.yaml 失败: %v", err)
		return nil, false
	}
	var dep DepConfig
	if err := yaml.Unmarshal(data, &dep); err != nil {
		failBuild(record, "解析 .dep/dep.yaml 失败: %v", err)
		return nil, false
	}
	appendLog(record, "代码拉取完成，解析 dep.yaml: %d 个步骤, %d 个产物", len(dep.Steps), len(dep.Artifacts))
	return &dep, true
}

// runSteps 按顺序执行编译步骤
func runSteps(record *models.BuildRecordModel, workDir string, dep *DepConfig) bool {
	for i, step := range dep.Steps {
		appendLog(record, "执行步骤[%d] %s: %s", i+1, step.Name, step.Cmd)
		out, err := utils.RunShell(workDir, buildTimeout(), step.Cmd)
		if out != "" {
			appendLog(record, "%s", strings.TrimSpace(out))
		}
		if err != nil {
			failBuild(record, "步骤[%d] %s 执行失败: %v", i+1, step.Name, err)
			return false
		}
	}
	return true
}

// uploadArtifacts 上传产物到七牛云，key 支持 {{version}} {{env}} {{project}} 占位
func uploadArtifacts(record *models.BuildRecordModel, project *models.ProjectModel, workDir string, dep *DepConfig, version string) bool {
	for _, artifact := range dep.Artifacts {
		localPath := filepath.Join(workDir, artifact.Path)
		key := renderPlaceholders(artifact.Key, project, record.Env, version)
		appendLog(record, "上传产物 %s -> %s", artifact.Path, key)
		url, err := utils.UploadQiniu(localPath, key)
		if err != nil {
			failBuild(record, "上传产物失败: %v", err)
			return false
		}
		appendLog(record, "上传成功: %s", url)
	}
	return true
}

// deployIfNeeded web 类工程且配置 deploy 时调用 k8s 发布
func deployIfNeeded(record *models.BuildRecordModel, project *models.ProjectModel, dep *DepConfig, version string) bool {
	if project.Type != models.ProjectTypeWeb || dep.Deploy.Name == "" {
		return true
	}
	if utils.Conf.K8s == "" {
		appendLog(record, "未配置 k8s 环境，跳过发布")
		return true
	}
	image := renderPlaceholders(dep.Deploy.Image, project, record.Env, version)
	appendLog(record, "发布到 k8s: deployment=%s container=%s image=%s", dep.Deploy.Name, dep.Deploy.Container, image)
	out, err := utils.K8sDeployImage(dep.Deploy.Name, dep.Deploy.Container, image)
	if out != "" {
		appendLog(record, "%s", strings.TrimSpace(out))
	}
	if err != nil {
		failBuild(record, "k8s 发布失败: %v", err)
		return false
	}
	return true
}

// renderPlaceholders 替换 {{version}} {{env}} {{project}} 占位符
func renderPlaceholders(s string, project *models.ProjectModel, env string, version string) string {
	s = strings.ReplaceAll(s, "{{version}}", version)
	s = strings.ReplaceAll(s, "{{env}}", env)
	s = strings.ReplaceAll(s, "{{project}}", project.Name)
	return s
}

// currentVersion 取工程当前环境版本号
func currentVersion(project *models.ProjectModel, env string) string {
	if env == models.BuildEnvProd {
		return project.ProdVersion
	}
	return project.TestVersion
}

// bumpVersion 版本号 patch 位 +1，无法解析时返回 1.0.1
func bumpVersion(version string) string {
	parts := strings.Split(version, ".")
	if len(parts) == 3 {
		if patch, err := strconv.Atoi(parts[2]); err == nil {
			return fmt.Sprintf("%s.%s.%d", parts[0], parts[1], patch+1)
		}
	}
	return "1.0.1"
}

// saveProjectVersion 构建成功后回写工程对应环境的版本号
func saveProjectVersion(project *models.ProjectModel, env string, version string) {
	field := "test_version"
	if env == models.BuildEnvProd {
		field = "prod_version"
	}
	if err := utils.DB.Model(project).Update(field, version).Error; err != nil {
		utils.Error("回写工程版本号失败 project=%d: %v", project.ID, err)
	}
}

func buildTimeout() time.Duration {
	return time.Duration(utils.Conf.Build.TimeoutMinutes) * time.Minute
}

func setStatus(record *models.BuildRecordModel, status string) {
	record.Status = status
	if err := utils.DB.Model(record).Update("status", status).Error; err != nil {
		utils.Error("更新构建状态失败 record=%d: %v", record.ID, err)
	}
}

func appendLog(record *models.BuildRecordModel, format string, args ...any) {
	line := time.Now().Format("2006-01-02 15:04:05") + " " + fmt.Sprintf(format, args...)
	record.Log += line + "\n"
	if err := utils.DB.Model(record).Update("log", record.Log).Error; err != nil {
		utils.Error("写入构建日志失败 record=%d: %v", record.ID, err)
	}
}

func failBuild(record *models.BuildRecordModel, format string, args ...any) {
	appendLog(record, "构建失败: "+format, args...)
	setStatus(record, models.BuildStatusFailed)
}
