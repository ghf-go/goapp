package action

import (
	"github.com/ghf-go/goapp/apppackage/app/job"
	"github.com/ghf-go/goapp/apppackage/app/models"
	"github.com/ghf-go/goapp/apppackage/app/utils"
	"github.com/gin-gonic/gin"
)

type buildStartReq struct {
	ProjectID uint64 `json:"projectId" binding:"required"`
	Env       string `json:"env" binding:"required"`
}

// BuildStartAction 启动一次构建：创建构建记录并异步执行
func BuildStartAction(c *gin.Context) {
	var req buildStartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "参数错误: "+err.Error())
		return
	}
	if req.Env != models.BuildEnvTest && req.Env != models.BuildEnvProd {
		utils.Fail(c, "环境不合法，可选: test/prod")
		return
	}
	var project models.ProjectModel
	if err := utils.DB.First(&project, req.ProjectID).Error; err != nil {
		utils.Fail(c, "工程不存在")
		return
	}
	branch := project.TestBranch
	if req.Env == models.BuildEnvProd {
		branch = project.ProdBranch
	}
	if branch == "" {
		utils.Fail(c, "该环境未配置分支")
		return
	}
	record := models.BuildRecordModel{
		ProjectID: project.ID,
		Env:       req.Env,
		Branch:    branch,
		Status:    models.BuildStatusPending,
	}
	if err := utils.DB.Create(&record).Error; err != nil {
		utils.Fail(c, "创建构建记录失败: "+err.Error())
		return
	}
	go job.RunBuild(record.ID)
	utils.Success(c, gin.H{"recordId": record.ID})
}
