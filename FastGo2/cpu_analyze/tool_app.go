package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func do_something(i int) {
	for j := 0; j < i; j++ {
		time.Sleep(time.Duration(i) * time.Second)
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", true, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", true, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Println("create cpu pprof failed,err:", err)
			return
		}
		/* go tool pprof .\cpu.pprof
		   flat：当前函数占用CPU的耗时
		   flat%：当前函数占用CPU耗时占总CPU耗时的百分比
		   sum%：函数占用CPU的耗时累计百分比
		   cum：当前函数加上当前函数调用函数占用CPU的总耗时
		   cum %：当前函数加上当前函数调用函数占用CPU总耗时百分比
		list 函数名，查看具体的函数分析
		*/

		err = pprof.StartCPUProfile(file)
		if err != nil {
			fmt.Println("start cpu pprof failed,err:", err)
			return
		}
		defer file.Close()
		defer pprof.StopCPUProfile()
	}

	// 业务代码
	for i := 1; i < 10; i++ {
		go do_something(i)
	}
	time.Sleep(5 * time.Second)

	if isMemPprof {
		file2, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Println("create mem pprof failed,err:", err)
			return
		}
		err = pprof.WriteHeapProfile(file2)
		if err != nil {
			fmt.Println("write heap pprof failed,err:", err)
			return
		}
		defer file2.Close()
	}
}
