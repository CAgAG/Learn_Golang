package AdvancedSort

import "DataSort/SimpleSort"

func CountSort(arr []int) []int {
	length := len(arr)
	max := SimpleSort.SelectSortMax(arr, length)
	sortedarr := make([]int, length)
	countarr := make([]int, max+1)

	for _, v := range arr {
		countarr[v]++
	}
	for i := 1; i <= max; i++ {
		countarr[i] += countarr[i-1]
	}
	for _, v := range arr {
		sortedarr[countarr[v]-1] = v
		countarr[v]--
	}
	return sortedarr
}
