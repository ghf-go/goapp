package help

import (
	"fmt"
	"strings"
)

func Print(msg string, tabs int) {
	fmt.Printf("%s%s\n", strings.Repeat("\t", tabs), msg)
}
