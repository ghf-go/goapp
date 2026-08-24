package action

import (
	"github.com/ghf-go/goapp/apppackage/app/models"
	"github.com/ghf-go/goapp/apppackage/app/utils"
	"github.com/gin-gonic/gin"
)

type idReq struct {
	ID uint64 `json:"id" binding:"required"`
}

// ProjectDeleteAction 删除工程
func ProjectDeleteAction(c *gin.Context) {
	var req idReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "参数错误: "+err.Error())
		return
	}
	if err := utils.DB.Delete(&models.ProjectModel{}, req.ID).Error; err != nil {
		utils.Fail(c, "删除失败: "+err.Error())
		return
	}
	utils.Success(c, nil)
}
