package AdvancedSort

import "math"

// 锦标赛排序

type node struct {
	value       int
	isAvailable bool // 无穷大(false)
	rank        int  // 叶子中的排序
}

func CompareAndUp(tree *[]node, leftnode int) {
	rightNode := leftnode + 1
	mid := (leftnode - 1) / 2
	if !(*tree)[leftnode].isAvailable ||
		((*tree)[rightNode].isAvailable && (*tree)[leftnode].value > (*tree)[rightNode].value) {
		(*tree)[mid] = (*tree)[rightNode]
	} else {
		(*tree)[mid] = (*tree)[leftnode]
	}
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func TreeSelectSort(arr []int) []int {
	length := len(arr)
	var (
		level  int
		result = make([]int, 0, length)
	)
	for Pow(2, level) < length {
		level++
	}
	var (
		leaf = Pow(2, level)
		tree = make([]node, leaf*2-1)
	)
	// 填充叶子节点
	for i := 0; i < length; i++ {
		tree[leaf+i-1] = node{arr[i], true, i}
	}
	// 对比
	// 每层都比较叶子兄弟大小，选出较大值作为父节点
	for i := 0; i < level; i++ {
		nodeCount := Pow(2, level-i) // 每次处理降低一个层级/2
		// 每组兄弟间比较
		for j := 0; j < nodeCount/2; j++ {
			leftnode := nodeCount - 1 + j*2
			CompareAndUp(&tree, leftnode)
		}
	}
	result = append(result, tree[0].value)
	// 选出一个之后, 还有n-1个
	for t := 0; t < length-1; t++ {
		winnode := tree[0].rank + leaf - 1 // 记录赢的节点
		tree[winnode].isAvailable = false  // 修改为无穷大, 无效
		// 从下一轮开始，只需与每次胜出节点的兄弟节点进行比较
		for i := 0; i < level; i++ {
			leftnode := winnode
			if winnode%2 == 0 { // 处理奇偶数
				leftnode = winnode - 1
			}
			// 比较兄弟节点间大小，并将胜出的节点向上传递
			CompareAndUp(&tree, leftnode)
			winnode = (leftnode - 1) / 2 // 保存中间节点
		}
		result = append(result, tree[0].value)
	}
	return result
}
