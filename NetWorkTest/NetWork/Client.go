package NetWork

import (
	"fmt"
	"net"
)

func Build() {
	addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8848")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}

	for true {
		var input string
		fmt.Scanln(&input)
		if input == "q"{
			break
		}

		conn.Write([]byte(input))
		buffer := make([]byte, 1024)
		n, _ := conn.Read(buffer)
		fmt.Println(string(buffer[: n]))
	}
	conn.Close()
}

