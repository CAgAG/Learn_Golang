package SimpleSort

import "DataStructure/SingleLink"

func InsertSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 1; i < length; i++ {
			backup := arr[i]
			j := i-1
			for j>=0 && backup<arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = backup
		}
	}
	return arr
}

func InsertSortLink(head *SingleLink.SingleLinkNode) *SingleLink.SingleLinkNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := &SingleLink.SingleLinkNode{0, nil}
	cur := head
	for cur != nil {
		next := cur.Next
		pre := newHead
		for pre.Next != nil && pre.Next.Value.(int) < cur.Value.(int) {
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		cur = next
	}
 	return newHead.Next
}
