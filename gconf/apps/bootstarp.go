package apps

import (
	"os"

	"github.com/ghf-go/goapp/gconf/apps/kimi"
	"github.com/ghf-go/goapp/gconf/help"
)

func Run() {
	if len(os.Args) < 2 {
		helpUsage()
		return
	}
	switch os.Args[1] {
	case "kimi":
		kimi.Run()
	default:
		helpUsage()
	}
}
func helpUsage() {
	help.Print("Usage:", 0)
	help.Print("kimi [command]", 1)
	help.Print("", 0)

}
