package BinSearch

func BinSearch(arr []int, data int) int {
	low := 0
	high := len(arr) -1
	for low <= high {
		mid := (low+high)/2
		if arr[mid] > data {
			high = mid-1
		} else if arr[mid] < data {
			low = mid+1
		} else {
			return mid
		}
	}
	return -1
}

// 第一个
func BinSearch_First(arr []int, data int) int {
	low := 0
	high := len(arr) -1
	index := -1

	for low <= high {
		mid := (low+high)/2
		if arr[mid] > data {
			high = mid-1
		} else if arr[mid] < data {
			low = mid+1
		} else {
			if mid==0 || arr[mid-1]!=data {
				index = mid
				break
			} else {
				high = mid - 1
			}
		}
	}
	return index
}

// 最后一个
func BinSearch_Last(arr []int, data int) int {
	low := 0
	high := len(arr) -1
	index := -1

	for low <= high {
		mid := (low+high)/2
		if arr[mid] > data {
			high = mid-1
		} else if arr[mid] < data {
			low = mid+1
		} else {
			if mid==len(arr)-1 || arr[mid+1]!=data {
				index = mid
				break
			} else {
				low = mid + 1
			}
		}
	}
	return index
}

// >=
func BinSearch_bigger(arr []int, data int) int {
	low := 0
	high := len(arr) -1
	index := -1

	for low <= high {
		mid := (low+high)/2
		if arr[mid] < data {
			low = mid+1
		} else {
			if mid==0 || arr[mid-1]<data{
				index = mid
				break
			} else {
				high = mid - 1
			}
		}
	}
	return index
}

// <=
func BinSearch_lesser(arr []int, data int) int {
	low := 0
	high := len(arr) -1
	index := -1

	for low <= high {
		mid := (low+high)/2
		if arr[mid] > data {
			high = mid-1
		} else {
			if mid==len(arr)-1 || arr[mid+1]>data{
				index = mid
				break
			} else {
				low = mid + 1
			}
		}
	}
	return index
}