package AdvancedSort

import "DataSort/SimpleSort"

// 用于重复值多的, 数据量有限
func BucketSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		MaxNum := SimpleSort.SelectSortMax(arr, length)
		buckets := make([][]int, MaxNum)
		for i := 0; i < length; i++ {
			buckets[arr[i]-1] = append(buckets[arr[i]-1], arr[i]) // 桶计数+1
		}
		tppose := 0
		for i := 0; i < MaxNum; i++ {
			bucketsLen := len(buckets[i])
			if bucketsLen > 0 {
				copy(arr[tppose: ], buckets[i])
				tppose += bucketsLen // 定位
			}
		}
	}
	return arr
}

func BucketSortX(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		num := length
		MaxNum := SimpleSort.SelectSortMax(arr, length)

		buckets := make([][]int, num)
		index := 0

		for i := 0; i < length; i++ {
			index = arr[i]*(num-1)/MaxNum // 木桶的自动分配算法, 会产生空桶
			buckets[index] = append(buckets[index], arr[i]) // 桶计数+1
		}
		tppose := 0
		for i := 0; i < num; i++ {
			bucketsLen := len(buckets[i])
			if bucketsLen > 0 {
				buckets[i] = SimpleSort.SelectSort(buckets[i]) // 木桶内部数据排序

				copy(arr[tppose: ], buckets[i])
				tppose += bucketsLen // 定位
			}
		}
	}
	return arr
}
