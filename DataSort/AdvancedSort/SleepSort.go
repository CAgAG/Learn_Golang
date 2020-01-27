package AdvancedSort

import (
	"fmt"
	"time"
)

// 用于分布式, 多线程
var container chan bool // 管道
var flag bool
var count int

func ToSleep(data int) {
	time.Sleep(time.Duration(data)*time.Microsecond*1000)
	fmt.Print(data, " ")
	container <- true // 往管道写入true
}

func Listen(size int) {
	for flag {
		select {
		case <-container:
			count++ // 计数器
			if count >= size { // 等待5个数值输入成功
				flag = false
				break
			}
		}

	}
}

func SleepSort(arr []int) {
	flag = true
	container = make(chan bool, 5) // 5个管道
	for i := 0; i < len(arr); i++ {
		go ToSleep(arr[i]) // 并发
	}
	go Listen(len(arr))
	for flag {
		time.Sleep(1*time.Second)
	}

}
