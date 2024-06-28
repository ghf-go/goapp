package action

import (
	"fmt"
	"runtime"

	"github.com/ghf-go/goapp/base"
)

// 添加包安装器
func brewAction() {
	fmt.Println(runtime.GOOS)
	switch runtime.GOOS {
	case "darwin":
		base.ShRun("xcode-select --install;git clone --depth=1 https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/install.git brew-install;/bin/bash brew-install/install.sh;rm -rf brew-install")
	case "linux":
	case "win":
	}
}
