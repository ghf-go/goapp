package base

import (
	"flag"
	"fmt"
	"os/exec"
)

var args = []string{}

//获取请求参数

func GetArg(i int) string {
	if len(args) == 0 {
		flag.Parse()
		args = flag.Args()
	}
	lenArgs := len(args)
	if i >= lenArgs {
		return ""
	}
	return args[i]
}

type ActionFunc func()

// 执行请求
func RunAction(i int, callMap map[string]ActionFunc) {
	cmd := GetArg(i)
	if call, ok := callMap[cmd]; ok {
		call()
	} else {
		cmd = "help"
		if call, ok := callMap[cmd]; ok {
			call()
		} else {
			fmt.Println("参数错误")
		}

	}
}

// 运行命令行
func ShRun(cmd string) {
	fmt.Sprintln(cmd)
	exec.Command("bash", "-c", cmd)
}
