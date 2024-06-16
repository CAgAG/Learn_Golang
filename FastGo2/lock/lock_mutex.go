package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁
// 只能被一个goroutine同时持有。如果另一个goroutine试图获取一个已被持有的互斥锁，它将被阻塞，直到持有锁的goroutine释放锁。
var lock_mutex sync.Mutex
var list = []int{}

func other_read_g() {
	lock_mutex.Lock()
	<-time.After(3 * time.Second)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("other read!")
	fmt.Println("other read: ", list[3])
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	lock_mutex.Unlock()
}

func other_write_g() {
	lock_mutex.Lock()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("other write!")
	list[3] = 10
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	lock_mutex.Unlock()
}

func main() {
	fmt.Println("互斥锁测试!!!")

	for i := 0; i < 10; i++ {
		list = append(list, i)
	}

	go other_read_g()
	go other_write_g()

	<-time.After(1 * time.Second)
	// 使用互斥锁保护共享资源
	lock_mutex.Lock()
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println("main read!")
	fmt.Println("main read: ", list[3])
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	lock_mutex.Unlock()
}
