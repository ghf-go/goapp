package utils

import (
	"fmt"
	"time"
)

// K8sDeployImage 调用本机 kubectl 更新 deployment 镜像
// deployment/container/image 均来自 dep.yaml 的 deploy 配置
func K8sDeployImage(deployment string, container string, image string) (string, error) {
	if deployment == "" || container == "" || image == "" {
		return "", fmt.Errorf("k8s发布参数不完整(deployment/container/image)")
	}
	cmd := fmt.Sprintf("kubectl set image deployment/%s %s=%s", deployment, container, image)
	return RunShell("", 2*time.Minute, cmd)
}
