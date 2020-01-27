package AdvancedSort

import "fmt"

// 表插入排序
type Node struct {
	Value int
	Next int // 下一个值的索引
}

const Int_MAX = int(^uint(0) >> 1) // 位操作

var NL []Node

func InitList(arr []int) {
	node := Node{Int_MAX, 1}
	NL = append(NL, node)
	for i := 1; i < len(arr); i++ {
		node = Node{arr[i-1], 0}
		NL = append(NL, node)
	}
	fmt.Println(NL)
}

func ListSort() {
	var (
		i, low, high int
	)
	for i = 2; i < len(NL); i++ {
		low = 0
		high = NL[0].Next
		for NL[high].Value < NL[i].Value { // 寻找一个邻居的数据 NL[max] NL[i]，插入NL[min]
			low = high
			high = NL[high].Next
		}
		NL[low].Next = i
		NL[i].Next = high // 插入数据到中间
	}
}

// 地址排序，插入排序
func Arrange() {
	p := NL[0].Next
	for i := 1; i < len(NL); i++ {
		for p < i {
			p = NL[p].Next
		}
		q := NL[p].Next
		if q != i {
			NL[p].Value, NL[i].Value = NL[i].Value, NL[p].Value
			NL[p].Next = NL[i].Next // 修改next
			NL[i].Next = p // 地址插入
		}
		p = q
	}
}


