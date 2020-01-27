package SimpleSort

func SelectSortMax(arr []int, length int) int {
	if length <= 1{
		return arr[0]
	} else {
		max := arr[0]
		for i := 0; i < length; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}