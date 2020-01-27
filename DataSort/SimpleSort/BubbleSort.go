package SimpleSort

import "DataStructure/SingleLink"

func BubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			needexchange := false
			for j := 0; j < length-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					needexchange = true
				}
			}
			if !needexchange{
				break
			}
		}
	}
	return arr
}

func BubbleSortLink(head *SingleLink.SingleLinkNode) {
	if head == nil || head.Next == nil {
		return
	}
	for p1 := head.Next; p1.Next != nil; p1 = p1.Next {
		for p2 := head.Next; p2.Next != nil; p2 = p2.Next {
			if p2.Value.(int) < p2.Next.Value.(int) {
				p2.SwapValue(p2.Next)
			}
		}
	}
}


