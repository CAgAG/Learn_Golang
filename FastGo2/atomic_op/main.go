package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var num int64 = 0

	fmt.Println(num)
	fmt.Println(atomic.AddInt64(&num, 2))
	fmt.Println(num)

	fmt.Println(atomic.AddInt64(&num, -4))
	fmt.Println(num)

	fmt.Println(atomic.CompareAndSwapInt64(&num, -2, 0)) // 旧(old)值 和 num的值相等才会替换
	fmt.Println(num)
	
}
