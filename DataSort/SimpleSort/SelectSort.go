package SimpleSort

import "DataStructure/SingleLink"

func SelectSortString(arr [] string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length; i++ {
			min := i
			for j := i + 1; j < length; j++ {
				if arr[min] > arr[j] {
					min = j
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i]
			}
		}
	}
	return arr
}

func SelectSort(arr [] int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length; i++ {
			min := i
			for j := i + 1; j < length; j++ {
				if arr[min] > arr[j] {
					min = j
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i]
			}
		}
	}
	return arr
}

func SelectSortLink(head *SingleLink.SingleLinkNode) {
	if head == nil || head.Next == nil {
		return
	}

	SortedNode := head.Next
	for SortedNode != nil {
		MaxNode := SortedNode
		p := SortedNode

		for p != nil {
			if p.Value.(int) > MaxNode.Value.(int) {
				MaxNode = p
			}
			p = p.Next
		}
		MaxNode.SwapValue(SortedNode)
		SortedNode = SortedNode.Next
	}



}