package AdvancedSort

import (
	"math/rand"
	"time"
)

// 是否已排好
func isOrder(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

// 洗牌算法
func RandList(arr []int) {
	length := len(arr)
	data := make([]int, length)
	copy(data, arr)
	rand.Seed(time.Now().UnixNano()) // 随机种子
	index := rand.Perm(length) // 产生指定长度的随机索引
	for i, k := range index {
		arr[i] = data[k]
	}
}

func RandSort(arr []int) []int {
	for true {
		if isOrder(arr) {
			break
		} else {
			RandList(arr)
		}
	}
	return arr
}
