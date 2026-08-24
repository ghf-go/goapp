package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一返回结构：{code, msg, data}，code=0 成功
const (
	CodeSuccess      = 0
	CodeFail         = 1
	CodeUnauthorized = 401
)

// Success 成功返回
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"code": CodeSuccess, "msg": "success", "data": data})
}

// Fail 失败返回
func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": CodeFail, "msg": msg, "data": nil})
}

// Unauthorized 未登录/登录过期返回
func Unauthorized(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": CodeUnauthorized, "msg": msg, "data": nil})
}
