package CircleLink

import "fmt"

func Add(node *Node) {
	if Tail == nil {
		Head = node
		node.Next = Head
		Tail = node
	} else {
		Tail.Next = node
		node.Next = Head
		Tail = node
	}
}

func Show(head *Node) {
	if head == nil {
		return
	} else {
		for head.Next != nil && head != Tail {
			fmt.Print(head.Data, " ")
			head = head.Next
		}
		fmt.Println(head.Data)
	}
}

// 约瑟夫环
// 从第K个，循环起第num个，留下最后一个
func Jose(k, num int) {
	count := 1 // 记录次数
	for i := 0; i <= k-1; i++ {
		Head = Head.Next
		Tail = Tail.Next // 循环到起点
	}
	for true {
		count++
		Head = Head.Next
		Tail = Tail.Next

		if count == num {
			Tail = Head.Next
			Head = Head.Next

			count = 1
		}
		if Head == Tail { // 相等意味着仅剩一个
			break
		}
	}
}
