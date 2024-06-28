package action

import (
	"fmt"
	"runtime"

	"github.com/ghf-go/goapp/base"
)

func Run() {
	fmt.Println(runtime.GOOS, runtime.GOARCH, runtime.GOROOT())
	base.RunAction(0, map[string]base.ActionFunc{
		"help":   helpAction,
		"brew":   brewAction,
		"nodejs": nodeJsAction,
	})
}
