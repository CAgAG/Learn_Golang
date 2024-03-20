package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
goroutine 奉行通过通信来共享内存，而不是共享内存来通信。
channel是Go语言中的一个核心类型，可以把它看成管道。并发核心单元通过它就可以发送或者接收数据进行通讯，这在一定程度上又进一步降低了编程的难度。
*/

func newTask() {
	i := 0
	for i < 3 {
		i++
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second) //延时1s
	}
	fmt.Println("stop cur goroutine 1")
	runtime.Goexit()                    // 终止当前 goroutine, import "runtime"
	fmt.Println("stop cur goroutine 2") // 不会执行
}

func main4() {
	// 创建一个 goroutine，启动另外一个任务
	// 只需在函数调⽤语句前添加 go 关键字，就可创建并发执⾏单元。开发⼈员无需了解任何执⾏细节，调度器会自动将其安排到合适的系统线程上执行。
	go newTask()
	i := 0
	//main goroutine 循环打印
	for i < 5 {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second) //延时1s
	}
}
