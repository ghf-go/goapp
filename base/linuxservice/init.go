package linuxservice

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/ghf-go/goapp/base/gpath"
)

func Run() bool {
	isRun := false
	for _, item := range os.Args {
		val := strings.ToLower(strings.TrimSpace(item))
		switch val {
		case "-installservice":
			isRun = true
			InstallService()
		case "-uninstallservice":
			isRun = true
			UnInstallService()
		}
	}
	return isRun
}

// 输出	Init 系统	service 支持情况
// systemd	systemd	自带 service 兼容层，完全支持
// init	SysVinit	原生支持 service
// upstart	Upstart	原生支持 service
// openrc	OpenRC(Alpine)	无 service，使用 rc-service

// 获取设备支持的类型
func getInitSystem() string {
	cmd := exec.Command("ps", "-p", "1", "-o", "comm=")
	var out bytes.Buffer
	cmd.Stdout = &out
	// 执行命令
	if err := cmd.Run(); err != nil {
		return ""
	}
	// 去除换行、空格
	name := strings.TrimSpace(out.String())
	return name
}

// 添加service
func InstallService() {
	switch getInitSystem() {
	case "systemd":
		installSystemd(gpath.GetSelfPath())
	case "openrc":
		break
	}
}

// 删除service
func UnInstallService() {
	switch getInitSystem() {
	case "systemd":
		uninstallSystemd(gpath.GetSelfPath())
	case "openrc":
		break
	}
}
