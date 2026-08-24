package action

import (
	"strconv"

	"github.com/ghf-go/goapp/apppackage/app/models"
	"github.com/ghf-go/goapp/apppackage/app/utils"
	"github.com/gin-gonic/gin"
)

// ProjectListAction 工程分页列表，支持名称模糊搜索
func ProjectListAction(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 10
	}
	query := utils.DB.Model(&models.ProjectModel{})
	if name := c.Query("name"); name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		utils.Fail(c, "查询失败: "+err.Error())
		return
	}
	var list []models.ProjectModel
	err := query.Order("id DESC").Offset((page - 1) * size).Limit(size).Find(&list).Error
	if err != nil {
		utils.Fail(c, "查询失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{"list": list, "total": total})
}
