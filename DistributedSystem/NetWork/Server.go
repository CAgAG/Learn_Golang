package NetWork

import (
	"log"
	"net"
	"os/exec"
	"time"
)

// 超30s未产生通信, 则退出
func HeartBeat(conn net.Conn, heartchan chan byte, timeout int) {
	select {
	case hc := <-heartchan:
		log.Println("heartchan: ", string(hc))
		conn.SetDeadline(time.Now().Add(time.Duration(timeout)*time.Second))
	case <-time.After(time.Second * 30):
		log.Println("time out: ", conn.RemoteAddr()) // 客户端超时
		conn.Close()
	}
}

func HeartChanHandler(n []byte, beatch chan byte) {
	for _, v := range n {
		beatch <- v
	}
	close(beatch)
}

func MsgHandler(conn net.Conn) {
	buffer := make([]byte, 1024)
	defer conn.Close()

	for true {
		n, err := conn.Read(buffer)
		if err != nil{
			return
		}
		msg := buffer[1: n]
		if n != 0{
			if string(buffer[0:1]) == "0" {
				conn.Write([]byte("收到数据: "+string(buffer[1: n]) + "\n"))
			} else {
				cmd := exec.Command(string(buffer[1: n])) // 执行命令
				cmd.Run()
				conn.Write([]byte("收到命令: "+string(buffer[1:n])+"\n"))
			}
		}

		beatch := make(chan byte)
		go HeartBeat(conn, beatch, 30)
		go HeartChanHandler(msg, beatch)
	}
}





