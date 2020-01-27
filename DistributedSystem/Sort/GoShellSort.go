package Sort

import (
	"runtime"
	"sync"
)

//性能调优，希尔排序改造多线程
func ShellSortGoRoutine(arr []int) {
	length := len(arr)
	if length < 2 || arr == nil {
		return
	}
	wg := sync.WaitGroup{} // 等待多个线程返回
	GoRoutineNum := runtime.NumCPU() //抓取系统有几个CPU

	for gap := length / 2; gap > 0; gap /= 2 {
		wg.Add(GoRoutineNum)
		ch := make(chan int, 10000) // 通道，进行线程通信
		go func() {
			for k := 0; k < gap; k++ {
				ch <- k // 输入到通道中
			}
			close(ch) // 关闭通道
		}()
		for k := 0; k < GoRoutineNum; k++ {
			go func() {
				for v := range ch{
					ShellSortStep(arr, v, gap)
				}
				wg.Done() // 一直等待完成
			}()
		}
		wg.Wait()
	}

}

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start; i < length; i += gap {
		back := arr[i]
		j := i-gap
		for j >= 0 && back < arr[j] {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = back
	}
}