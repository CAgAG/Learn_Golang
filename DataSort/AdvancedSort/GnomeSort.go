package AdvancedSort

// 处理几乎有序的情况, 已经快要排序完的数据
func GnomeSort(arr []int) []int {
	i := 1
	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			if i > 1 {
				i--
			}
		}
	}
	return arr
}
