package app

import (
	"fmt"
	"os"

	"github.com/ghf-go/goapp/gtools/app/models"
	"github.com/ghf-go/goapp/gtools/app/utils"
)

var (
	_models map[string]models.BaseModel
)

func Run() {
	initModel()
	argsLen := len(os.Args)
	if argsLen < 2 {
		help()
		os.Exit(1)
	} else {
		getModelAndRun(os.Args[1])
	}
}

// 初始化model
func initModel() {
	_models = map[string]models.BaseModel{

		"login":    models.NewLoginModel(),
		"autocode": models.NewAiAgentSkillRuleModel(),
	}
}

// 并运行脚本
func getModelAndRun(name string) {
	m, ok := _models[name]
	if !ok {
		help()
		os.Exit(0)
	}
	args := utils.BuildArgs()
	if e := m.Run(args); e != nil {
		fmt.Printf("run error: %s\n", e.Error())
		os.Exit(0)
	}
}
func help() {
	fmt.Printf("%s command\n\ncommans:\n", os.Args[0])
	for k, m := range _models {
		fmt.Printf("  %s  -> %s\n", k, m.GetDesc())
		fmt.Printf("    %s\n\n", m.GetUsage())
	}
}
