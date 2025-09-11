package main

import (
	"fmt"
	"log"
	"net"
)

// GOOS=linux GOARCH=arm64 go build -o ipscanserver main.go
const (
	serverAddress = ":9631" // 监听所有网络接口的8888端口
	response      = "ok"
	bufferSize    = 1024
)

func main() {
	// 监听UDP端口
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 9631})
	if err != nil {
		log.Fatalf("无法监听UDP端口: %v", err)
	}
	defer conn.Close()

	log.Printf("UDP广播服务端已启动，监听端口 %s", serverAddress)
	log.Println("等待接收广播消息...")
	for {
		// 接收消息
		buffer := make([]byte, bufferSize)
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("接收消息失败: %v", err)
			continue
		}
		log.Printf("收到来自 %s 的消息: %s", clientAddr.String(), string(buffer[:n]))
		sAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("255.255.255.255:%d", clientAddr.Port))
		// 发送响应
		n, err = conn.WriteToUDP([]byte(response), sAddr)
		if err != nil {
			log.Printf("发送响应失败: %v", err)
			continue
		}
		log.Printf("发送响应: %v:%v", n, response)
	}
}
