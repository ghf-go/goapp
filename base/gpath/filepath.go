package gpath

import (
	"fmt"
	"os"
	"path/filepath"
)

// 获取应用程序的绝对路径
func GetSelfPath() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println("程序完整路径:", exePath)

	// 3. 消除软链接，得到真实物理路径
	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		panic(err)
	}
	return realPath
}
