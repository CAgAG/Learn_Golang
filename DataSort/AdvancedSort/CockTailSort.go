package AdvancedSort
// 鸡尾酒排序-双向冒泡

func CockTailSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		left := 0
		right := length-1
		for left <= right {
			if arr[left] > arr[left+1] {
				arr[left], arr[left+1] = arr[left+1], arr[left]
			}
			left++
			if arr[right-1] > arr[right] {
				arr[right-1], arr[right] = arr[right], arr[right-1]
			}
			right--
		}
	}
	return arr
}
