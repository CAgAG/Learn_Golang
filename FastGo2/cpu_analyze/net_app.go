package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

func main() {
	// go tool pprof http://127.0.0.1:8080/debug/pprof/profile
	// go tool pprof http://127.0.0.1:8080/debug/pprof/profile?seconds=10

	// # 下载heap profile
	// go tool pprof http://127.0.0.1:8080/debug/pprof/heap
	http.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Println(req.Form)
		fmt.Println("path:", req.URL.Path)
		fmt.Println("scheme:", req.URL.Scheme)
		for k, v := range req.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
		resp.Write([]byte("hello world"))
	})

	// 如果你使用的是gin框架，那么推荐使用github.com/gin-contrib/pprof，在代码中通过以下命令注册pprof相关路由。
	// pprof.Register(router)

	fmt.Println(fmt.Sprintf("http://127.0.0.1:8080/hello"))
	fmt.Println(fmt.Sprintf("http://127.0.0.1:8080/debug/pprof/"))
	http.ListenAndServe("127.0.0.1:8080", nil)
}
