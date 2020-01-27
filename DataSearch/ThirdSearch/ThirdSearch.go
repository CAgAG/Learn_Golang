package ThirdSearch

// 分布式
func ThirdSearch(arr []int, data int) int {
	low := 0
	high := len(arr)-1

	for low <= high {
		mid1 := low + (high-low)/3
		mid2 := high - (high-low)/3
		mid1_data := arr[mid1]
		mid2_data := arr[mid2]
		if mid1_data == data {
			return mid1
		} else if mid2_data == data {
			return mid2
		}

		if mid1_data < data {
			low = mid1 + 1
		} else if mid2_data > data {
			high = mid2 - 1
		} else {
			low = low + 1
			high = high - 1
		}
	}
	return -1
}
