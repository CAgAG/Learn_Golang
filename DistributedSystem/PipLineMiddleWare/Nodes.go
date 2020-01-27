package PipLineMiddleWare

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var StartTime time.Time // 时间

// 开始计时
func Init() {
	StartTime = time.Now()
}

// 消耗的时间
func UseTime() {
	fmt.Println(time.Since(StartTime))
}

// 内存排序
func InMemorySort(in <-chan int) <-chan int{
	out := make(chan int, 1024)
	go func() {
		data := []int{}
		for v := range in {
			data = append(data, v)
		}
		fmt.Println("data Read: ", UseTime)
		sort.Ints(data) // 排序
		for _, v := range data {
			out <- v
		}
		close(out)
	}()
	return out
}

// 合并通道
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				// 取出V1，压入，再次读取v1
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}

// 读取数据
func ReadSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8) // 64
		readsize := 0
		for true {
			n, err := reader.Read(buffer)
			readsize += n
			if n > 0 {
				out <- int(binary.BigEndian.Uint64(buffer)) // 数据压入
			}
			if err != nil || (chunkSize != -1 && readsize >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

// 写入
func WriteSlink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v)) // 字节转换
		writer.Write(buffer)
	}
}

// 随机产生数组
func RandSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

// 多路合并
func MergeN(inputs... <-chan int) <-chan int {
	length := len(inputs)
	if length == 1 {
		return inputs[0]
	}
	m := length/2
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}

func ArraySource(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range num {
			out <- v
		}
		close(out)
	}()
	return out
}


