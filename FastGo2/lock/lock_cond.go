package main

import (
	"fmt"
	"sync"
	"time"
)

// 一次性锁
// 是一个同步原语，它确保某个操作只被执行一次。
var lock_cond sync.Cond
var mu sync.Mutex
var list5 = []int{}

func other_write_g5(i int) {
	<-time.After(time.Duration(i) * time.Second)
	mu.Lock()
	lock_cond.Wait() // 等待信号

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("other write! ==> ", i)
	list5[3] = i
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	mu.Unlock()
}

func main() {
	fmt.Println("条件变量测试!!!")
	lock_cond.L = &mu

	for i := 0; i < 10; i++ {
		list5 = append(list5, i)
	}

	go other_write_g5(1)

	<-time.After(2 * time.Second)
	// 发送信号
	mu.Lock()
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println("main read!")
	lock_cond.Signal() // 发送信号
	mu.Unlock()

	<-time.After(1 * time.Second)

	fmt.Println("main read: ", list5[3])
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
}
