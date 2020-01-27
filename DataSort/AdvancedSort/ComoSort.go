package AdvancedSort

// ShellSort gap=1.3
func ComoSort(arr []int) []int {
	length := len(arr)
	gap := length
	for gap > 1 {
		gap = gap*10/13
		for i := 0; i+gap < length; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
			}
		}
	}
	return arr
}
