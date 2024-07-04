package main

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/ghf-go/goapp/base"
	"github.com/ghf-go/goapp/gocordova/action"
)

//go:embed template
var tdir embed.FS

// 获取当前路径
func getPwd() string {
	r, _ := os.Getwd()
	return r
}
func help() {
	fmt.Println("gocordova 使用帮助\n")
	fmt.Println("\tcreate name packageName // 创建一个应用 name:应用名称 packageName:包名")
	fmt.Println("\twxappimg src //src 为图片路径")
	fmt.Println("\tiosIcon src //src 为图片路径")
	fmt.Println("")
}

// 保存文件
func saveFiles(pf, p2, appName, appPackage, appDir string) {

	flist, e := tdir.ReadDir(pf)
	if e != nil {
		panic(e.Error())
	}
	for _, item := range flist {
		if !item.IsDir() {
			dd, e := tdir.ReadFile(pf + "/" + item.Name())
			if e != nil {
				panic(e.Error())
			}
			d1 := strings.ReplaceAll(string(dd), "myApp", appName)
			d2 := strings.ReplaceAll(d1, "org.apache.cordova.myApp", appPackage)
			dd = []byte(d2)
			if item.Name() == "MainActivity.java" {
				oldPath := strings.ReplaceAll("org.apache.cordova.myApp", ".", "/")
				newPath := strings.ReplaceAll(appPackage, ".", "/")
				p3 := strings.Replace(p2, oldPath, newPath, 1)
				os.RemoveAll(appDir + strings.Replace(p2, oldPath, "org", 1))
				os.MkdirAll(appDir+p3, 0777)
				e = os.WriteFile(appDir+p3+"/"+item.Name(), dd, 0644)
				if e != nil {
					panic(appDir + p2 + "/" + item.Name() + e.Error())
				}
				continue
			}
			if item.Name() == "gitignore" {
				e = os.WriteFile(appDir+p2+"/."+item.Name(), dd, 0644)
				if e != nil {
					panic(appDir + p2 + "/" + item.Name() + e.Error())
				}
			} else {
				e = os.WriteFile(appDir+p2+"/"+item.Name(), dd, 0644)
				if e != nil {
					panic(appDir + p2 + "/" + item.Name() + e.Error())
				}
			}

		} else {
			os.Mkdir(appDir+p2+"/"+item.Name(), 0777)
			saveFiles(pf+"/"+item.Name(), p2+"/"+item.Name(), appName, appPackage, appDir)
		}

	}
}

// 创建应用
func createApp() {

	name := base.GetArg(1)
	pname := base.GetArg(2)
	if name == "" {
		help()
		return
	}
	if pname == "" {
		help()
		return
	}
	appDir := getPwd() + "/" + name
	os.Mkdir(appDir, 0777)
	// dd, e := tdir.ReadFile("template/.gitignore")
	// if e != nil {
	// 	panic(e.Error() + string(dd))
	// 	//os.WriteFile(appDir+"/.gitignore", dd, 0644)
	// }

	saveFiles("template", "", name, pname, appDir)
}
func main() {
	base.RunAction(0, map[string]base.ActionFunc{
		"help":     help,
		"create":   createApp,
		"wxappimg": action.WxAppImgAction,
		"iosIcon":  action.IosIconAction,
	})
}
