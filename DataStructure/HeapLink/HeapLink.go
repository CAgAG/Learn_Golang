package HeapLink

import "fmt"

func NewLeftHeap(data interface{}) PQ {
	head := new(TreeNode)
	head.data = data
	head.left = nil
	head.right = nil
	head.npl = 0
	return PQ(head)
}

func MergeSort(H1, H2 PQ) PQ {
	if H1.left == nil {
		H1.left = H2
	} else {
		H1.right = Merge(H1.right, H2)
		if H1.left.npl < H1.right.npl {
			H1.left, H1.right = H1.right, H1.left
		}
		H1.npl = H1.right.npl+1 // 级递增 层
	}

	return H1
}


func Merge(H1, H2 PQ) PQ {
	if H1 == nil {
		return H2
	}
	if H2 == nil {
		return H1
	}
	if H1.data.(int) > H2.data.(int) {
		return MergeSort(H1, H2) // 确保左边小于右边
	} else {
		return MergeSort(H2, H1)
	}
}

func Insert(data interface{}, H PQ) PQ {
	node := new(TreeNode)
	node.data = data
	node.left = nil
	node.right = nil
	node.npl = 0

	H = Merge(node, H)
	return H
}

func DeleteMax(H PQ) (PQ, interface{}) {
	if H == nil {
		return nil, nil
	} else {
		left := H.left
		right := H.right
		v := H.data
		H = nil
		return Merge(left, right), v
	}
}

func PrintHeap(H PQ) {
	if H == nil {
		return
	}
	PrintHeap(H.left)
	PrintHeap(H.right)
	fmt.Print(H.data, " ")
}
















