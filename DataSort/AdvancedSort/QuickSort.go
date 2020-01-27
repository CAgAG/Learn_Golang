package AdvancedSort

import (
	"DataStructure/SingleLink"
	"math/rand"
)

func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		splitdata := arr[0] // 第一个为基准值
		low := make([]int, 0, 0)
		high := make([]int, 0, 0)
		mid := make([]int,0 ,0)
		mid = append(mid, splitdata)

		for i := 1; i < length; i++ {
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)
		arr := append(append(low, mid...), high...)
		return arr
	}
}

func QuickSortX(arr []int) []int {
	length := len(arr)
	if length <= 1{
		return arr
	} else {
		n := rand.Int() % length
		splitdata := arr[n]
		low := make([]int, 0, 0)
		high := make([]int, 0, 0)
		mid := make([]int,0 ,0)
		mid = append(mid, splitdata)

		for i := 1; i < length; i++ {
			if i == n {
				continue
			}
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)
		arr := append(append(low, mid...), high...)
		return arr
	}
}

func GetPartionLink(left, right *SingleLink.SingleLinkNode) *SingleLink.SingleLinkNode {
	splitdata := left.Value.(int)
	p := left
	q := left.Next

	for q != right.Next {
		if q.Value.(int) < splitdata {
			p = p.Next
			p.SwapValue(q)
		}
		q = q.Next
	}
	p.SwapValue(left)
	return p
}

func QuickSortLinkImple(left, right *SingleLink.SingleLinkNode)  {
	if left != right {
		parition := GetPartionLink(left, right)
		QuickSortLinkImple(left, parition)
		if parition.Next != nil {
			QuickSortLinkImple(parition.Next, right)
		}
	}
}

func QuickSortLink(head *SingleLink.SingleLinkNode) *SingleLink.SingleLinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	right := head
	for right.Next != nil {
		right = right.Next
	}
	QuickSortLinkImple(head, right)
	return head
}
