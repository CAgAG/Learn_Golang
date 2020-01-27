package FindNLargest

func Swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func Partition(arr []int, left, right int) int {
	pivot := right

	i := left
	for j := left; j < pivot; j++ {
		if arr[j] > arr[pivot] {
			Swap(arr, i, j)
			i++
		}
	}
	//最后交换arr[pivot]和arr[i]的位置, 则可以确保arr[pivot]左边的元素都比自己小, 右边的元素都比自己大
	Swap(arr, i, pivot)
	return i
}

func QuickSortImple(arr []int, left, right int) {
	if left >= right {
		return
	}
	q := Partition(arr, left, right)
	QuickSortImple(arr, left, q-1)
	QuickSortImple(arr, q+1, right)
}

func QuickSort(arr []int) []int {
	QuickSortImple(arr, 0, len(arr)-1)
	return arr
}

func FindNLargestImple(arr []int, left, right, n int) int {
	if left >= right {
		return arr[left]
	}
	query := Partition(arr, left, right)
	if query+1 == n {
		return arr[query]
	}
	if n < query+1 {
		return FindNLargestImple(arr, left, query-1, n)
	}
	return FindNLargestImple(arr, query+1, right, n)
}

func FindNLargest(arr []int, n int) int {
	return FindNLargestImple(arr, 0, len(arr)-1, n)
}