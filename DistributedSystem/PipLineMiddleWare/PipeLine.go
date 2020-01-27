package PipLineMiddleWare

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func DealErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 分布式
func CreateNetWordPipeLine(fileName string, fileSize int, chunkCount int) <-chan int {
	file, err := os.Create(fileName)
	DealErr(err)

	defer file.Close()
	Pipe := RandSource(fileSize/8)
	writer := bufio.NewWriter(file)
	WriteSlink(writer, Pipe)
	writer.Flush() // 刷新

	chunkSize := fileSize/chunkCount
	sortAddr := []string{} // 分布式地址
	Init()
	file, err = os.Open(fileName)
	DealErr(err)
	defer file.Close()

	for i := 0; i < chunkSize; i++ {
		file.Seek(int64(i*chunkSize), 0) // 移动文件指针位置
		source := ReadSource(bufio.NewReader(file), chunkSize)
		addr := ":" + strconv.Itoa(7000 + i) // 开辟地址

		NetWordWrite(addr, InMemorySort(source)) // 写入到分布式主机
		sortAddr = append(sortAddr, addr)
	}
	sortResult := []<-chan int{}
	for _, addr := range sortAddr {
		sortResult = append(sortResult, NetWordRead(addr))
	}
	return MergeN(sortResult...)
}

// 单机多线程
func CreatePipeLine(fileName string, fileSize int, chunkCount int) <-chan int {
	file, err := os.Create(fileName)
	DealErr(err)

	defer file.Close()
	Pipe := RandSource(fileSize/8)
	writer := bufio.NewWriter(file)
	WriteSlink(writer, Pipe)
	writer.Flush() // 刷新

	chunkSize := fileSize/chunkCount
	sortResult := []<-chan int{}
	Init()
	file, err = os.Open(fileName)
	DealErr(err)
	defer file.Close()

	for i := 0; i < chunkSize; i++ {
		file.Seek(int64(i*chunkSize), 0) // 移动文件指针位置
		source := ReadSource(bufio.NewReader(file), chunkSize)
		sortResult = append(sortResult, InMemorySort(source))
	}

	return MergeN(sortResult...)
}

// 写入文件
func Write2File(in <-chan int, fileName string) {
	file, err := os.Create(fileName)
	DealErr(err)

	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	WriteSlink(writer, in)
}

// 最多显示文件前1000行内容
func ShowFile(fileName string) {
	file, err := os.Open(fileName)
	DealErr(err)

	defer file.Close()
	p := ReadSource(bufio.NewReader(file), -1)

	counter := 0
	for v := range p {
		fmt.Println(v)
		counter++
		if counter > 1000 {
			break
		}
	}
}
