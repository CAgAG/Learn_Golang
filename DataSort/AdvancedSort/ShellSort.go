package AdvancedSort

func ShellSortStep(arr []int, start int, gap int, length int) {
	for i := start + gap; i < length; i += gap { // 插入排序变化
		backup := arr[i]
		j := i-gap
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup
	}
}

func ShellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		gap := length/2
		for gap > 0 {
			for i := 0; i < gap; i++ {
				ShellSortStep(arr, i, gap, length)
			}
			gap /= 2
		}
	}
	return arr
}
