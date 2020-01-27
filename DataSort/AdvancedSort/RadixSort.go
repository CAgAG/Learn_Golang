package AdvancedSort

import "DataSort/SimpleSort"

func BiSort(arr []int, bit int, length int) []int {
	bitcounts := make([]int, 10)
	for i := 0; i < length; i++ {
		num := (arr[i]/bit)%10
		bitcounts[num]++
	}
	for i := 1; i < 10; i++ {
		bitcounts[i] += bitcounts[i-1]
	}
	tp := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		num := (arr[i]/bit)%10
		tp[bitcounts[num]-1] = arr[i]
		bitcounts[num]--
	}
	for i := 0; i < length; i++ {
		arr[i] = tp[i]
	}
	return arr
}

func RadixSort(arr []int) []int {
	length := len(arr)
	max := SimpleSort.SelectSortMax(arr, length)
	for bit := 1; max/bit > 0; bit *= 10 {
		arr = BiSort(arr, bit, length)
	}
	return arr
}
