package action

import (
	"strconv"

	"github.com/ghf-go/goapp/apppackage/app/models"
	"github.com/ghf-go/goapp/apppackage/app/utils"
	"github.com/gin-gonic/gin"
)

// BuildRecordDetailAction 构建记录详情（含完整日志）
func BuildRecordDetailAction(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Query("id"), 10, 64)
	if id == 0 {
		utils.Fail(c, "参数错误: id 必填")
		return
	}
	var record models.BuildRecordModel
	if err := utils.DB.First(&record, id).Error; err != nil {
		utils.Fail(c, "构建记录不存在")
		return
	}
	var project models.ProjectModel
	projectName := ""
	if err := utils.DB.First(&project, record.ProjectID).Error; err == nil {
		projectName = project.Name
	}
	utils.Success(c, gin.H{
		"id":          record.ID,
		"projectId":   record.ProjectID,
		"projectName": projectName,
		"env":         record.Env,
		"branch":      record.Branch,
		"version":     record.Version,
		"status":      record.Status,
		"log":         record.Log,
		"createdAt":   record.CreatedAt,
		"updatedAt":   record.UpdatedAt,
	})
}
