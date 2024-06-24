package action

import (
	"fmt"
	"os"
	"runtime"

	"github.com/ghf-go/goapp/base"
)

func Run() {
	os.Setenv("ghf", "123")
	fmt.Println(runtime.GOOS, runtime.GOARCH, runtime.GOROOT())
	base.RunAction(0, map[string]base.ActionFunc{
		"help":  helpAction,
		"mac":   macAction,
		"linux": linuxAction,
	})
}
