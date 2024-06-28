package action

import (
	"runtime"

	"github.com/ghf-go/goapp/base"
)

func nodeJsAction() {
	switch runtime.GOOS {
	case "darwin":
		base.ShRun("brew install nodejs npm;npm config set registry https://registry.npmmirror.com/;npm install -g nvm;nvm install stable;npm install -g @vue/cli;npm install -g cordova")
	case "linux":
	case "win":
	}
}
