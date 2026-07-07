package gpath

import (
	"path/filepath"
	"strings"
)

// 获取不带文件后缀的文件名
func GetFileNameWithoutExt(path string) string {
	// 取出文件名（含后缀）
	fileName := filepath.Base(path)
	// 取出后缀 .xxx
	ext := filepath.Ext(fileName)
	// 去除后缀
	nameNoExt := strings.TrimSuffix(fileName, ext)
	return nameNoExt
}
