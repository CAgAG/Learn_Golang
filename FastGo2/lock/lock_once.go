package main

import (
	"fmt"
	"sync"
	"time"
)

// 一次性锁
// 是一个同步原语，它确保某个操作只被执行一次。
var lock_once sync.Once
var list4 = []int{}

func other_write_g4(i int) {
	<-time.After(time.Duration(i) * time.Second)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("other write! ==> ", i)

	lock_once.Do(func() {
		list4[3] = i
	})
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
}

func main() {
	fmt.Println("一次性锁测试!!!")

	for i := 0; i < 10; i++ {
		list4 = append(list4, i)
	}

	go other_write_g4(1)
	go other_write_g4(2)

	// 等待
	<-time.After(5 * time.Second)
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println("main read!")
	fmt.Println("main read: ", list4[3])
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")

}
