package BinSearch

import "fmt"

func MakeFabArr(arr []int) []int {
	// 创建fab Array
	length := len(arr)
	flen := 2
	f, s, t := 1, 2, 3
	for t < length {
		t, f, s = f+s, s, t
		flen++
	}
	fbs := make([]int, flen)
	fmt.Println(flen)
	fbs[0] = 1
	fbs[1] = 1
	for i := 2; i < flen; i++ {
		fbs[i] = fbs[i-1]+fbs[i-2]
	}
	return fbs
}

func Fab_Search(arr []int, val int) int {
	length := len(arr)
	fabArr := MakeFabArr(arr)
	fillLen := fabArr[len(fabArr)-1]

	fillArr := make([]int, fillLen)
	for i, v := range arr{
		fillArr[i] = v
	}

	lastData := arr[length-1] // 最大数替代0
	for i := length; i < fillLen; i++ {
		fillArr[i] = lastData
	}

	left, mid, right := 0, 0, length
	kindex := len(fabArr)-1

	for left <= right {
		mid = left+fabArr[kindex-1]-1 // fab 切割
		if val < fillArr[mid]{
			right = mid - 1
			kindex--
		} else if val > fillArr[mid] {
			left = mid + 1
			kindex -= 2
		} else {
			if mid > right {
				return right
			} else {
				return mid
			}
		}

	}

	return -1
}
