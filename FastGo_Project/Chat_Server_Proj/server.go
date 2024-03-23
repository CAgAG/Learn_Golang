// 服务端
package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	IP   string
	PORT int

	// 在线用户列表
	OnlineMap map[string]*User
	map_lock  sync.RWMutex // 同步锁

	// 全局消息广播
	Message chan string
}

// 创建
func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:        ip,
		PORT:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message，一旦有变动就发送给全体用户
func (this *Server) ListenMessage() {
	for {
		msg, ok := <-this.Message
		if !ok {
			break
		} else {
			this.map_lock.Lock()
			for _, user := range this.OnlineMap {
				user.Channel <- msg
			}
			this.map_lock.Unlock()
		}
	}
}

// 广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMessage := fmt.Sprintf("[%s]%s: %s", user.Addr, user.Name, msg)

	this.Message <- sendMessage
}

func (this *Server) Handler(conn net.Conn) {
	// 业务代码
	user := NewUser(conn, this)
	defer func() {
		close(user.Channel)
		conn.Close()
	}()
	fmt.Println("链接建立成功:", user.Addr)

	// 用户上线
	user.Online()
	// 监听用户是否活跃
	is_live := make(chan bool)

	// 接收用户发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				// 用户下线
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn error:", err)
				return
			}
			// 提取用户信息, 同时去除最后一位 note: linux:\n windows: \r\n
			msg := string(buf[:n-1])
			// 用户发送信息
			user.DoMessage(msg)

			// 更新活跃信息
			is_live <- true
		}
	}()

	// 阻塞
	for true {
		select {
		// 用户是否活跃
		case <-is_live:
			// 活跃就不操作
		// 60秒不发言就踢出
		case <-time.After(time.Second * 60):
			user.SelfMessage("You have been removed from the group chat!")
			return
		}
	}
}

func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.PORT))
	if err != nil {
		fmt.Println("net.Listen error: ", err)
		return
	}
	// close server
	defer listener.Close()

	// 开始监听全局Message
	go this.ListenMessage()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error: ", err)
			continue
		}
		// handler
		go this.Handler(conn)
	}

}
