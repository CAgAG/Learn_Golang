package AdvancedSort
// 奇偶排序

func OddEven(arr []int) []int {
	isSort := false
	length := len(arr)
	for isSort == false {
		isSort = true
		for i := 1; i < length-1; i+=2 { // 奇数
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSort = false
			}
		}
		for i := 0; i < length-1; i += 2 { // 偶数
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSort = false
			}
		}
	}
	return arr
}
