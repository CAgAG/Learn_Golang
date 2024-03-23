package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	// 用户消息
	Channel chan string
	Conn    net.Conn
	server  *Server
}

// 创建用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:    userAddr,
		Addr:    userAddr,
		Channel: make(chan string),
		Conn:    conn,
		server:  server,
	}

	// 开始监听消息
	go user.ListenMessage()

	return user
}

// 监听当前user，一旦有消息加入就立刻写入返回客户端
func (this *User) ListenMessage() {
	for {
		msg, ok := <-this.Channel
		if !ok { // 管道如果已经关闭就退出循环
			break
		} else {
			this.Conn.Write([]byte(msg + "\n"))
		}
	}
}

// 用户操作
// 上线
func (this *User) Online() {
	// 用户上线
	// 用户加入到在线表
	this.server.map_lock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.map_lock.Unlock()
	// 全体广播上线信息
	this.server.BroadCast(this, "Online")
}

// 下线
func (this *User) Offline() {
	delete(this.server.OnlineMap, this.Name)
	this.server.BroadCast(this, "Offline")

	defer this.Conn.Close()
}

// 给当前用户自己发送消息
func (this *User) SelfMessage(msg string) {
	this.Conn.Write([]byte(msg + "\n"))
}

// 处理消息
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		this.server.map_lock.Lock()
		for user_name, _ := range this.server.OnlineMap {
			online_msg := fmt.Sprintf("[%s]%s: Online", this.Addr, user_name)
			this.SelfMessage(online_msg)
		}
		this.server.map_lock.Unlock()
	} else if len(msg) > 7 && strings.TrimSpace(msg)[:6] == "rename" {
		// 重命名用户名称
		new_name := strings.Split(msg, " ")[1]
		// 判断name是否存在
		_, ok := this.server.OnlineMap[new_name]
		if ok {
			this.SelfMessage(fmt.Sprintf("%s already exist\n", new_name))
		} else {
			this.server.map_lock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[new_name] = this
			this.server.map_lock.Unlock()

			this.Name = new_name
			this.SelfMessage(fmt.Sprintf("new user name: %s\n", new_name))
		}
	} else if len(msg) > 4 && strings.TrimSpace(msg)[:2] == "to" {
		// to user_name sentences
		if len(strings.Split(strings.TrimSpace(msg), " ")) < 3 {
			this.SelfMessage("A wrong 'to' command!")
		}
		// 接收的用户名称
		receive_name := strings.Split(msg, " ")[1]
		// 获取 user对象
		to_user, ok := this.server.OnlineMap[receive_name]
		if !ok {
			this.SelfMessage("There is not a user name: " + receive_name)
		} else {
			to_msg := strings.Join(strings.Split(msg, " ")[2:], " ")
			to_user.SelfMessage(fmt.Sprintf("From %s: %s", this.Name, to_msg))
		}
	} else {
		this.server.BroadCast(this, msg)
	}
}
