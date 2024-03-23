package main

import (
	"flag"
	"fmt"
)

var (
	global_server_ip   string
	global_server_port int
)

func init() {
	// 命令行解析
	flag.StringVar(&global_server_ip, "ip", "127.0.0.1", "set server ip")
	flag.IntVar(&global_server_port, "port", 8888, "set server port")
}

func main() {
	/*
		go build -o client.exe .\client.go .\main.go
		.\client.exe -h
		  -ip string
		        set server ip (default "127.0.0.1")
		  -port int
		        set server port (default 8888)
	*/
	flag.Parse()

	client := NewClient(global_server_ip, global_server_port)
	if client == nil {
		fmt.Println("xxx Failed to connect to server! xxx")
		return
	}

	fmt.Println("*** The server connection is successful! ***")

	// 单独开启一个goroutine去处理server给client的消息
	go client.DealResponse()
	client.Start()
}
