package action

import (
	"github.com/ghf-go/goapp/apppackage/app/utils"
	"github.com/gin-gonic/gin"
)

type loginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginAction 登录，校验配置中的管理员账号
func LoginAction(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "参数错误: "+err.Error())
		return
	}
	if req.Username != utils.Conf.Admin.Username || req.Password != utils.Conf.Admin.Password {
		utils.Fail(c, "账号或密码错误")
		return
	}
	utils.Success(c, gin.H{"token": utils.GenerateToken()})
}

// LogoutAction 退出登录
func LogoutAction(c *gin.Context) {
	utils.RemoveToken(c.GetHeader("Authorization"))
	utils.Success(c, nil)
}

// ProfileAction 返回当前登录账号信息
func ProfileAction(c *gin.Context) {
	utils.Success(c, gin.H{"username": utils.Conf.Admin.Username})
}
