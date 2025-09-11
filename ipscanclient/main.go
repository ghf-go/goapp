package main

import (
	"fmt"
	"net"
	"time"
)

const (
	broadcastAddress = "255.255.255.255:9631"
	message          = "Hello, Broadcast!"
	timeout          = 20 * time.Second
	bufferSize       = 2
)

func main() {
	// 服务器地址（目标地址）
	serverAddr, err := net.ResolveUDPAddr("udp", broadcastAddress)

	if err != nil {
		fmt.Printf("解析地址失败: %v\n", err)
		return
	}
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 1369})
	if err != nil {
		fmt.Printf("无法创建UDP连接: %v\n", err)
	}
	defer conn.Close()

	// 设置读写超时
	err = conn.SetDeadline(time.Now().Add(timeout))
	if err != nil {
		fmt.Printf("设置超时失败: %v\n", err)
	}
	go func() {
		// 发送广播消息
		_, err = conn.WriteToUDP([]byte(message), serverAddr)
		if err != nil {
			fmt.Printf("发送广播消息失败: %v\n", err)
		}
	}()

	// 接收响应
	buffer := make([]byte, bufferSize)
	_, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Printf("等待响应超时 %v -> %v\n", err, addr)
			return
		}
	}

	// 提取并显示服务器IP地址
	serverIP, _, err := net.SplitHostPort(addr.String())
	if err != nil {
		fmt.Printf("解析服务器地址失败: %v\n", err)
		serverIP = addr.String() // 失败时使用原始地址
	}

	fmt.Printf("可用服务器: %v\n", serverIP)
}
