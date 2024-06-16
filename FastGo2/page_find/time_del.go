package main

import (
	"fmt"
	"sync"
	"time"
)

type TempArray struct {
	arr     []int
	mutex   sync.Mutex
	timeout time.Duration
}

func NewTempArray(timeout time.Duration) *TempArray {
	return &TempArray{
		arr:     make([]int, 0),
		mutex:   sync.Mutex{},
		timeout: timeout,
	}
}

func (ta *TempArray) Add(value int) {
	ta.mutex.Lock() // 加锁
	ta.arr = append(ta.arr, value)
	ta.mutex.Unlock() // 解锁

	// 设置定时器，在timeout时间后删除数组中的一个元素
	timer := time.NewTimer(ta.timeout)
	go func() {
		<-timer.C
		ta.mutex.Lock() // 加锁
		ta.Remove(value)
		ta.mutex.Unlock() // 解锁
		fmt.Println(fmt.Sprintf("remove value: %d", value))
	}()
}

func (ta *TempArray) Remove(value int) {
	for i, v := range ta.arr {
		if v == value {
			ta.arr = append(ta.arr[:i], ta.arr[i+1:]...)
			break
		}
	}
}

func (ta *TempArray) Print() {
	ta.mutex.Lock() // 加锁
	fmt.Println("当前数组：", ta.arr)
	ta.mutex.Unlock() // 解锁
}

func main() {
	tempArray := NewTempArray(2 * time.Second)
	tempArray.Add(1)
	tempArray.Add(2)
	tempArray.Add(3)
	tempArray.Print()

	time.Sleep(5 * time.Second)
	tempArray.Print()
}
