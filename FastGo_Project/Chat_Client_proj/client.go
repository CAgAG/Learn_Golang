package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
	mode       int // 模式
}

func NewClient(IP string, Port int) *Client {
	client := &Client{
		ServerIP:   IP,
		ServerPort: Port,
		Name:       fmt.Sprintf("%s:%d", IP, Port),
		mode:       -1,
	}
	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", IP, Port))
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return nil
	}
	client.conn = conn
	return client
}

// 处理server回应的消息，直接显示到标准输出即可
func (this *Client) DealResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准输出上，永久阻塞监昕
	io.Copy(os.Stdout, this.conn)
}

func (this *Client) Rename() bool {
	fmt.Println("new user name: ")
	var new_name string
	fmt.Scanln(&new_name)
	this.Name = strings.TrimSpace(new_name)

	command := fmt.Sprintf("rename %s\n", this.Name)
	_, err := this.conn.Write([]byte(command))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

func (this *Client) Menu() bool {
	fmt.Println("************************")
	fmt.Println("1 global chat")
	fmt.Println("2 single user chat")
	fmt.Println("3 change user name")
	fmt.Println("0 exit")
	fmt.Println("************************")

	var flag_str string
	fmt.Scanln(&flag_str)
	flag, err := strconv.Atoi(flag_str)
	if err != nil {
		fmt.Println("xxx A wrong command number! xxx")
		return false
	}

	if flag >= 0 && flag <= 3 {
		// 合法命令输入
		this.mode = flag
		return true
	} else {
		fmt.Println("xxx A wrong command number! xxx")
		return false
	}
}

func (this *Client) PublicChat() {
	var msg string = ""
	fmt.Println("Input 'exit' to return to the menu")

	for {
		fmt.Scanln(&msg)
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			return
		}

		msg = fmt.Sprintf("%s\n", msg)
		_, err := this.conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}

		msg = ""
		fmt.Println("Input 'exit' to return to the menu")
	}
}

func (this *Client) SelectAllUser() {
	msg := "who\n"
	fmt.Println("************************")
	_, err := this.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	fmt.Println("************************")
}

func (this *Client) UserChat() {
	var to_user string
	var msg string
	this.SelectAllUser()

	for {
		fmt.Println("Input chat user, and input 'exit' to return to the menu:")
		fmt.Scanln(&to_user)
		to_user = strings.TrimSpace(to_user)
		if to_user == "exit" || to_user == "" {
			break
		} else {
			for {
				fmt.Println("Input message, and input 'exit' to return to the menu:")
				fmt.Scanln(&msg)
				msg = strings.TrimSpace(msg)

				if msg == "exit" {
					break
				} else if msg == "" {
					continue
				} else {
					msg = fmt.Sprintf("to %s %s\n", to_user, msg)
					_, err := this.conn.Write([]byte(msg))
					if err != nil {
						fmt.Println("conn.Write err:", err)
						break
					}
				}

				msg = ""
			}
		}
		to_user = ""
	}
}

func (this *Client) Start() {
	for this.mode != 0 {
		for this.Menu() != true {
			// 循环要求合法命令输入
		}
		switch this.mode {
		case 1:
			// 公聊模式
			fmt.Println("*** global chat ***")
			this.PublicChat()
			break
		case 2:
			// 私聊模式
			fmt.Println("*** single user chat ***")
			this.UserChat()
			break
		case 3:
			// 重命名
			fmt.Println("*** change user name ***")
			this.Rename()
			break
		case 0:
			// 结束
			fmt.Println("*** exit ***")
			break
		}
	}
}
