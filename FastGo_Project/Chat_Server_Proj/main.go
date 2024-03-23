package main

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}

// windows
// curl.exe --http0.9 127.0.0.1:8888

// linux
// nc 127.0.0.1 8888
