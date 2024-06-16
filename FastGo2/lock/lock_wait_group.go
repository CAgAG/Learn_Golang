package main

import (
	"fmt"
	"sync"
	"time"
)

// 等待组
// 是一个同步原语，它允许一个goroutine等待其他一组goroutine完成。
var lock_wg sync.WaitGroup
var list3 = []int{}

func other_read_g3() {
	lock_wg.Add(1) // 添加一个等待组

	<-time.After(1 * time.Second)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("other read!")
	fmt.Println("other read: ", list3[3])
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	lock_wg.Done()
}

func other_write_g3(i int) {
	lock_wg.Add(1) // 添加一个等待组

	<-time.After(time.Duration(i) * time.Second)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println("other write! ==> ", i)
	list3[3] = 10
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	lock_wg.Done()
}

func main() {
	fmt.Println("等待组测试!!!")

	for i := 0; i < 10; i++ {
		list3 = append(list3, i)
	}

	go other_read_g3()
	go other_write_g3(2)

	// 等待
	<-time.After(1 * time.Second)
	lock_wg.Wait()
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println("main read!")
	fmt.Println("main read: ", list3[3])
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")

}
