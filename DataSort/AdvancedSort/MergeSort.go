package AdvancedSort

import (
	"DataSort/SimpleSort"
	"DataStructure/SingleLink"
)

func Merge(left []int, right []int) []int {
	leftindex := 0
	rightindex := 0
	lastarr := make([]int, 0, 0)
	for leftindex < len(left) && rightindex < len(right) {
		if left[leftindex] < right[rightindex] {
			lastarr = append(lastarr, left[leftindex])
			leftindex++
		} else if left[leftindex] > right[rightindex] {
			lastarr = append(lastarr, right[rightindex])
			rightindex++
		} else {
			lastarr = append(lastarr, left[leftindex])
			lastarr = append(lastarr, right[rightindex])
			leftindex++
			rightindex++
		}
	}
	for leftindex < len(left) {
		lastarr = append(lastarr, left[leftindex])
		leftindex++
	}
	for rightindex < len(right) {
		lastarr = append(lastarr, right[rightindex])
		rightindex++
	}
	return lastarr
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else if length > 1 && length < 5 {
		return SimpleSort.InsertSort(arr)
	} else {
		mid := length/2
		left := MergeSort(arr[: mid])
		right := MergeSort(arr[mid: ])
		return Merge(left, right)
	}
}

func MergeLink(left, right *SingleLink.SingleLinkNode,) *SingleLink.SingleLinkNode {
	newNode := &SingleLink.SingleLinkNode{0, nil}
	p := newNode
	for left != nil && right != nil {
		if left.Value.(int) < right.Value.(int) {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
	}
	if left != nil { p.Next = left }
	if right != nil { p.Next = right }
	return newNode.Next
}

func MergeSortLink(head *SingleLink.SingleLinkNode) *SingleLink.SingleLinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	faster := head
	lower := head
	// 中间节点
	var mid *SingleLink.SingleLinkNode
	for faster != nil && faster.Next != nil {
		mid = lower
		lower = lower.Next
		faster = faster.Next.Next
	}
	mid.Next = nil
	left := MergeSortLink(head)
	rigth := MergeSortLink(lower)
	return MergeLink(left, rigth)
}
