package main

/*
   channel <- value      //发送value到channel
   <-channel             //接收并将其丢弃
   x := <-channel        //从channel中接收数据，并赋值给x
   x, ok := <-channel    //功能同上，同时检查通道是否已关闭或者是否为空
   =========================================================
   var ch1 chan int       // ch1是一个正常的channel，是双向的
   var ch2 chan<- float64 // ch2是单向channel，只用于写float64数据
   var ch3 <-chan int     // ch3是单向channel，只用于读int数据
   =========================================================
	Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。
    select {
    case <- chan1:
        // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
        // 如果成功向chan2写入数据，则进行该case处理语句
    default:
        // 如果上面都没有成功，则进入default处理流程
    }
*/

import (
	"fmt"
	"time"
)

//	chan<- //只写
//
// 生产者
func counter(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		fmt.Println("生产: ", i)
		out <- i //如果对方不读 会阻塞
	}
}

//	<-chan //只读
//
// 消费者
func printer(in <-chan int) {
	for num := range in {
		fmt.Println("消费: ", num)
	}
}

func main5() {
	c := make(chan int, 2)
	c2 := make(chan int, 0) //创建无缓冲的通道 c2

	go func() {
		defer fmt.Println("子go程结束")

		fmt.Println("子go程正在运行……")

		for i := 0; i < 5; i++ {
			fmt.Printf("c  子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
			c <- i
		}
		// 无缓冲的通道 c2
		for i := 0; i < 5; i++ {
			fmt.Printf("c2 子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c2), cap(c2))
			c2 <- i
		}

		// 关闭 Channel
		if _, ok := <-c; ok {
			close(c)
		}
		// 关闭 Channel
		if _, ok := <-c2; ok {
			close(c2)
		}
	}()

	for i := 0; i < 5; i++ {
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("c  main num = ", num)
	}

	fmt.Println("=====================================")
	time.Sleep(2 * time.Second) //延时2s
	for i := 0; i < 5; i++ {
		num := <-c2 //从c中接收数据，并赋值给num
		fmt.Println("c2 main num = ", num)
	}
	fmt.Println("main go程结束")

	fmt.Println("=====================================")
	c3 := make(chan int) //   chan   //读写
	go counter(c3)       //生产者
	printer(c3)          //消费者

	fmt.Println("done")
}
