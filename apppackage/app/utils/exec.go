package utils

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

// RunShell 在指定目录下通过 sh -c 执行命令，带超时，返回合并输出
func RunShell(dir string, timeout time.Duration, command string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	cmd.Dir = dir
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err := cmd.Run()
	out := buf.String()
	if ctx.Err() == context.DeadlineExceeded {
		return out, fmt.Errorf("命令执行超时(%s): %s", timeout, command)
	}
	if err != nil {
		return out, fmt.Errorf("命令执行失败: %w\n%s", err, out)
	}
	return out, nil
}
