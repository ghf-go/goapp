package gpath

import (
	"fmt"
	"os"
	"path/filepath"
)

// 获取应用程序的绝对路径
func GetSerlfPath() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println("程序完整路径:", exePath)

	// 2. 获取程序所在目录（去掉文件名）
	exeDir := filepath.Dir(exePath)
	fmt.Println("程序所在目录:", exeDir)

	// 3. 消除软链接，得到真实物理路径
	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		panic(err)
	}
	realDir := filepath.Dir(realPath)
	return realDir
}
