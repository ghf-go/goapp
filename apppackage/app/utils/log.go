package utils

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

// Log 输出普通日志
func Log(format string, args ...any) {
	logger.Printf("[INFO] "+format, args...)
}

// Error 输出错误日志
func Error(format string, args ...any) {
	logger.Printf("[ERROR] "+format, args...)
}
