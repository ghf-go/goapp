package kimi

import (
	"embed"
	"os"

	"github.com/ghf-go/goapp/gconf/help"
)

//go:embed agents/go.md agents/mobile.md agents/ros2.md agents/wails.md agents/platformio.md
var goAgentMdFs embed.FS

func Run() {
	if len(os.Args) < 3 {
		helpUsage()
		return
	}
	switch os.Args[2] {
	case "agents":
		agents()
	default:
		helpUsage()
	}
}
func agents() {
	if len(os.Args) < 4 {
		helpUsage()
		return
	}
	switch os.Args[3] {
	case "go":
		agentsRun("go")
	case "mobile":
		agentsRun("mobile")
	case "ros2":
		agentsRun("ros2")
	case "wails":
		agentsRun("wails")
	case "platformio":
		agentsRun("platformio")
	default:
		helpUsage()
	}
}
func agentsRun(name string) {
	data, e := goAgentMdFs.ReadFile("agents/" + name + ".md")
	if e != nil {
		helpUsage()
		return
	}
	if os.WriteFile("AGENTS.md", data, os.ModePerm) != nil {
		helpUsage()
		return
	}
}

func helpUsage() {
	help.Print("Usage:", 0)
	help.Print("kimi agents go 		#创建go项目的AGENTS.md", 1)
	help.Print("kimi agents mobile  	#创建移动项目的AGENTS.md", 1)
	help.Print("kimi agents ros2 	#创建ROS项目的AGENTS.md", 1)
	help.Print("kimi agents wails 	#创建wails项目的AGENTS.md", 1)
	help.Print("kimi agents platformio 	#创建platformio项目的AGENTS.md", 1)
	help.Print("", 0)
}
