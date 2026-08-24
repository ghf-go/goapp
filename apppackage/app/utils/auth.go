package utils

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	tokenStore   = map[string]time.Time{}
	tokenStoreMu sync.Mutex
)

// GenerateToken 生成登录 token，有效期取配置 admin.tokenExpireHours
func GenerateToken() string {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		panic("生成token失败: " + err.Error())
	}
	token := hex.EncodeToString(buf)
	tokenStoreMu.Lock()
	tokenStore[token] = time.Now().Add(time.Duration(Conf.Admin.TokenExpireHours) * time.Hour)
	tokenStoreMu.Unlock()
	return token
}

// CheckToken 校验 token 是否有效
func CheckToken(token string) bool {
	tokenStoreMu.Lock()
	defer tokenStoreMu.Unlock()
	expire, ok := tokenStore[token]
	if !ok {
		return false
	}
	if time.Now().After(expire) {
		delete(tokenStore, token)
		return false
	}
	return true
}

// RemoveToken 退出登录时移除 token
func RemoveToken(token string) {
	tokenStoreMu.Lock()
	delete(tokenStore, token)
	tokenStoreMu.Unlock()
}

// AuthMiddleware gin 中间件：校验 Authorization 头中的 token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !CheckToken(token) {
			Unauthorized(c, "未登录或登录已过期")
			return
		}
		c.Next()
	}
}
