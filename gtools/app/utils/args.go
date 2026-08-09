package utils

import (
	"os"
	"strings"
)

type Args struct {
	args []string
	kv   map[string]string
}

func BuildArgs() *Args {
	ret := &Args{
		args: os.Args[2:],
		kv:   make(map[string]string),
	}
	ret.parse()
	return ret
}
func (a *Args) Get(key string) string {
	r, ok := a.kv[key]
	if !ok {
		return ""
	}
	return r
}
func (a *Args) GetAll() []string {
	return a.args
}
func (a *Args) GetIndex(i int) string {
	if i >= len(a.args) {
		return ""
	}
	return a.args[i]
}
func (a *Args) parse() {
	for i, arg := range a.args {
		if strings.HasPrefix(arg, "-") {
			kk := strings.Trim(arg, "-")
			a.kv[kk] = a.GetIndex(i + 1)
		}
	}
}
