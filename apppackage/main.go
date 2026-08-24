package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/ghf-go/goapp/apppackage/app/action"
	"github.com/ghf-go/goapp/apppackage/app/utils"
	"github.com/gin-gonic/gin"
)

//go:embed conf
var confFS embed.FS

//go:embed app/views/dist
var distFS embed.FS

func main() {
	utils.InitConfig(confFS)
	utils.InitDB()
	startServer()
}

func startServer() {
	if !utils.Conf.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	registerAPI(r)
	registerStatic(r)
	addr := fmt.Sprintf("%s:%d", utils.Conf.Server.Host, utils.Conf.Server.Port)
	utils.Log("服务启动 %s", addr)
	if err := r.Run(addr); err != nil {
		panic("服务启动失败: " + err.Error())
	}
}

func registerAPI(r *gin.Engine) {
	r.POST("/api/login", action.LoginAction)
	api := r.Group("/api", utils.AuthMiddleware())
	{
		api.POST("/logout", action.LogoutAction)
		api.GET("/profile", action.ProfileAction)
		api.GET("/project/list", action.ProjectListAction)
		api.POST("/project/save", action.ProjectSaveAction)
		api.POST("/project/delete", action.ProjectDeleteAction)
		api.POST("/build/start", action.BuildStartAction)
		api.GET("/build/record/list", action.BuildRecordListAction)
		api.GET("/build/record/detail", action.BuildRecordDetailAction)
	}
}

// registerStatic 将 embed 的前端 dist 挂载到根路径，未命中路由回退到 index.html
func registerStatic(r *gin.Engine) {
	sub, err := fs.Sub(distFS, "app/views/dist")
	if err != nil {
		panic("加载前端资源失败: " + err.Error())
	}
	fileServer := http.FileServer(http.FS(sub))
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if len(path) >= 5 && path[:5] == "/api/" {
			c.JSON(http.StatusNotFound, gin.H{"code": utils.CodeFail, "msg": "接口不存在", "data": nil})
			return
		}
		if _, err := fs.Stat(sub, path[1:]); err != nil {
			c.Request.URL.Path = "/"
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	})
}
