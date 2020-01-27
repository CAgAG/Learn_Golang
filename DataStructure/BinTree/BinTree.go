package BinTree

import (
	"DataStructure/QueueArray"
	"DataStructure/StackArray"
	"fmt"
)

type BinaryTree struct {
	Root *Node
	Size int
}

func NewTree() *BinaryTree {
	bst := &BinaryTree{}
	bst.Root = nil
	bst.Size = 0
	return bst
}

func (bst *BinaryTree) GetSize() int {
	return bst.Size
}

func (bst *BinaryTree) IsEmpty() bool {
	return bst.Size == 0
}

func (bst *BinaryTree) Add(node *Node, data int) *Node {
	if node == nil {
		bst.Size += 1
		return &Node{data, nil, nil}
	}

	if data < node.Data {
		node.Left = bst.Add(node.Left, data)
	} else if data > node.Data {
		node.Right = bst.Add(node.Right, data)
	}
	return node
}

func (bst *BinaryTree) IsExist(node *Node, data int) bool {
	if node == nil {
		return false
	}
	if data == node.Data {
		return true
	} else if data < node.Data {
		return bst.IsExist(node.Left, data)
	} else {
		return bst.IsExist(node.Right, data)
	}

}

func (bst *BinaryTree) IsExistData(data int) bool {
	return bst.IsExist(bst.Root, data)
}

func (bst *BinaryTree) FindMaxByNode(node *Node) *Node {
	if node.Right == nil {
		return node
	} else {
		return bst.FindMaxByNode(node.Right)
	}
}

func (bst *BinaryTree) FindMax() int {
	if bst.Size == 0 {
		panic("empty tree")
	}
	return bst.FindMaxByNode(bst.Root).Data
}

func (bst *BinaryTree) FindMinByNode(node *Node) *Node {
	if node.Right == nil {
		return node
	} else {
		return bst.FindMaxByNode(node.Left)
	}
}

func (bst *BinaryTree) FindMin() int {
	if bst.Size == 0 {
		panic("empty tree")
	}
	return bst.FindMaxByNode(bst.Root).Data
}

func (bst *BinaryTree) PreOrderByNode(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Data, " ")
	bst.PreOrderByNode(node.Left)
	bst.PreOrderByNode(node.Right)
}

func (bst *BinaryTree) PreOrder() {
	bst.PreOrderByNode(bst.Root)
}

func (bst *BinaryTree) PreOrderStack() []int {
	StackType := StackArray.NewStack()

	bckTree := bst.Root   //备份
	stack := StackType    // 生成栈
	res := make([]int, 0) // 生成数组，容纳前序的数据

	for bckTree != nil || stack.Size() != 0 {

		for bckTree != nil {
			res = append(res, bckTree.Data) // 对应 fmt.Print(node.Data, " ")
			stack.Push(bckTree)
			bckTree = bckTree.Left // 首先左边压入栈中----- 对应 bst.PreOrderByNode(node.Left)
		}
		if stack.Size() != 0 {
			v := stack.Top()
			bckTree = v.(*Node)
			// ... res ?
			bckTree = bckTree.Right // 追加 ----- 对应 bst.PreOrderByNode(node.Right)
			stack.Pop()
		}
	}

	return res
}

func (bst *BinaryTree) InOrderByNode(node *Node) {
	if bst.Root == nil {
		return
	}
	bst.InOrder(node.Left)
	fmt.Print(node.Data, " ")
	bst.InOrder(node.Right)
}

func (bst *BinaryTree) InOrder(node *Node) {
	bst.InOrderByNode(bst.Root)
}

func (bst *BinaryTree) PostOrderByNode(node *Node) {
	if bst.Root == nil {
		return
	}
	bst.InOrder(node.Left)
	bst.InOrder(node.Right)
	fmt.Print(node.Data, " ")
}

func (bst *BinaryTree) PostOrder() {
	bst.PostOrderByNode(bst.Root)
}

//删除最小
func (bst *BinaryTree) RemoveMin() int {
	ret := bst.FindMin()
	bst.Root = bst.RemoveMinByNode(bst.Root)
	return ret
}
func (bst *BinaryTree) RemoveMinByNode(n *Node) *Node {
	if n.Left == nil {
		//删除
		rightNode := n.Right //备份右边节点
		bst.Size--           //删除
		return rightNode
	}
	n.Left = bst.RemoveMinByNode(n.Left)
	return n
}

//删除最大
func (bst *BinaryTree) RemoveMax() int {
	ret := bst.FindMax()
	bst.Root = bst.RemoveMaxByNode(bst.Root)
	return ret
}
func (bst *BinaryTree) RemoveMaxByNode(n *Node) *Node {
	if n.Right == nil {
		//删除
		leftNode := n.Left //备份左边节点
		bst.Size--         //删除
		return leftNode
	}
	n.Right = bst.RemoveMaxByNode(n.Right)
	return n
}

func (bst *BinaryTree) RemoveByNode(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if data < node.Data {
		node.Left = bst.RemoveByNode(node.Left, data)
		return node
	} else if data > node.Data {
		node.Right = bst.RemoveByNode(node.Right, data)
		return node
	} else {
		if node.Left == nil {
			rightNode := node.Right // 备份右边节点
			node.Right = nil
			bst.Size--
			return rightNode
		}
		if node.Right == nil {
			leftNode := node.Left // 备份左边节点
			node.Left = nil
			bst.Size--
			return leftNode
		}
		// 左右节点都不为空
		newNode := bst.FindMinByNode(node.Right)        // 找到小节点
		newNode.Right = bst.RemoveMinByNode(node.Right) // 删除小节点
		newNode.Left = node.Left

		node.Left = nil
		node.Right = nil
		return newNode
	}
}

func (bst *BinaryTree) GetDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}
	leftLen := bst.GetDepth(root.Left)
	rightLen := bst.GetDepth(root.Right)
	if leftLen > rightLen {
		return leftLen + 1
	} else {
		return rightLen + 1
	}
}

func (bst *BinaryTree) FindAncestor(root, a, b *Node) *Node {
	if root == nil {
		return nil
	}
	if root == a || root == b {
		return root
	}
	left := bst.FindAncestor(root.Left, a, b)
	right := bst.FindAncestor(root.Right, a, b)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}

// 层遍历
func (bst *BinaryTree) LevelShowByNode(node *Node) {
	QueueType := QueueArray.NewQueue()
	queue := QueueType

	queue.EnQueue(node)
	for queue.Size() > 0 {
		left := queue.Front()
		right := left
		queue.DeQueue()
		if v := right.(*Node); v != nil {
			fmt.Print(v.Data, " ")
			queue.EnQueue(v.Left)
			queue.EnQueue(v.Right)
		}
	}
}

func (bst *BinaryTree) LevelShow() {
	bst.LevelShowByNode(bst.Root)
}
