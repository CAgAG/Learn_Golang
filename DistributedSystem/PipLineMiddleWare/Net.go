package PipLineMiddleWare

import (
	"bufio"
	"net"
)

// 给地址(127.0.0.1:8090), 写入数据
func NetWordWrite(addr string, in <-chan int) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		// defer后面的函数, 在defer语句所在的函数执行结束的时候会被调用
		defer listen.Close() // 关闭网络
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		writer := bufio.NewWriter(conn) // 写入数据
		defer  writer.Flush()
		WriteSlink(writer, in)
	}()
}

// 给端口, 读取数据
func NetWordRead(addr string) <-chan int{
	out := make(chan int)
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		read := ReadSource(bufio.NewReader(conn), -1)
		for v := range read{
			out <- v
		}
		close(out)
	}()
	return out
}


