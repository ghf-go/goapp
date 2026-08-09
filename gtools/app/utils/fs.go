package utils

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

func IsExist(pathName string) bool {
	_, err := os.Stat(pathName)
	if err == nil {
		return true
	}
	// 文件不存在
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	// 其他错误（权限不足等）
	return false
}
func GetHomeDir() (string, error) {
	// 不同系统环境变量
	home := os.Getenv("HOME")
	if home == "" {
		// Windows
		home = os.Getenv("USERPROFILE")
	}
	if home != "" {
		return home, nil
	}

	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return u.HomeDir, nil
}

// 保存文件，如果目录不存在，则创建目录
func SaveFile(pathName string, data string) error {
	pdir := filepath.Dir(pathName)

	if !IsExist(pdir) {
		err := os.MkdirAll(pdir, 0755)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(pathName, []byte(data), 0644)
}

func BuildHomePathAndCreate(pathName string) string {
	home, err := GetHomeDir()
	if err != nil {
		return ""
	}
	fullPath := filepath.Join(home, pathName)

	pdir := filepath.Dir(fullPath)

	if !IsExist(pdir) {
		err = os.MkdirAll(pdir, 0755)
		if err != nil {
			return ""
		}
	}
	return fullPath
}
