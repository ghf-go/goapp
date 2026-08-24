package action

import (
	"strconv"

	"github.com/ghf-go/goapp/apppackage/app/models"
	"github.com/ghf-go/goapp/apppackage/app/utils"
	"github.com/gin-gonic/gin"
)

// BuildRecordListAction 构建记录分页列表，附带工程名称
func BuildRecordListAction(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 10
	}
	query := utils.DB.Model(&models.BuildRecordModel{})
	if projectID, _ := strconv.ParseUint(c.Query("projectId"), 10, 64); projectID > 0 {
		query = query.Where("project_id = ?", projectID)
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		utils.Fail(c, "查询失败: "+err.Error())
		return
	}
	var list []models.BuildRecordModel
	err := query.Order("id DESC").Offset((page - 1) * size).Limit(size).Find(&list).Error
	if err != nil {
		utils.Fail(c, "查询失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{"list": attachProjectName(list), "total": total})
}

// attachProjectName 为记录列表补充 projectName 字段
func attachProjectName(list []models.BuildRecordModel) []gin.H {
	ids := make([]uint64, 0, len(list))
	for _, r := range list {
		ids = append(ids, r.ProjectID)
	}
	var projects []models.ProjectModel
	nameMap := map[uint64]string{}
	if len(ids) > 0 {
		utils.DB.Where("id IN ?", ids).Find(&projects)
		for _, p := range projects {
			nameMap[p.ID] = p.Name
		}
	}
	result := make([]gin.H, 0, len(list))
	for _, r := range list {
		result = append(result, gin.H{
			"id":          r.ID,
			"projectId":   r.ProjectID,
			"projectName": nameMap[r.ProjectID],
			"env":         r.Env,
			"branch":      r.Branch,
			"version":     r.Version,
			"status":      r.Status,
			"createdAt":   r.CreatedAt,
			"updatedAt":   r.UpdatedAt,
		})
	}
	return result
}
