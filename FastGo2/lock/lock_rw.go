package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁
// 允许多个goroutine同时读取受保护的数据，但只允许一个goroutine同时写入受保护的数据。
var lock_rw sync.RWMutex
var list2 = []int{}

func other_read_g2() {
	<-time.After(1 * time.Second)
	lock_rw.RLock() // 读锁
	<-time.After(1 * time.Second)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("other read!")
	fmt.Println("other read: ", list2[3])
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	lock_rw.RUnlock()
}

func other_write_g2(i int) {
	lock_rw.Lock() // 写锁
	<-time.After(time.Duration(i) * time.Second)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("other write! ==> ", i)
	list2[3] = 10
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	lock_rw.Unlock()
}

func main() {
	fmt.Println("读写锁测试!!!")

	for i := 0; i < 10; i++ {
		list2 = append(list2, i)
	}

	go other_read_g2()
	go other_write_g2(1)
	go other_write_g2(2)

	<-time.After(2 * time.Second)
	// 使用互斥锁保护共享资源
	lock_rw.RLock() // 读锁
	<-time.After(1 * time.Second)
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println("main read!")
	fmt.Println("main read: ", list2[3])
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	lock_rw.RUnlock()
}
