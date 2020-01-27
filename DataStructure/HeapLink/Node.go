package HeapLink

type TreeNode struct {
	data interface{}
	left *TreeNode
	right *TreeNode
	npl int // 级别
}

type PQ *TreeNode // 优先队列
