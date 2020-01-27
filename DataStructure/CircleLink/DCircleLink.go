package CircleLink

import "fmt"

type DCircleLink struct {
	Id   int // 数据编号
	Data interface{}
	Prev *DCircleLink
	Next *DCircleLink
}

type DCImple interface {
	ResetHeadNode(data interface{})
	IsHeadEmpty() bool
	IsEmpty() bool
	GetLastNode() *DCircleLink
	AddNode(newNode *DCircleLink)
	FindNodeById(id int) (*DCircleLink, bool)
	DeleteNodeById(id int) bool
	ChangeNodeById(id int, data interface{}) bool
	Show()
}

func (head *DCircleLink) ResetHeadNode(data interface{}) {
	if head.Id == 0 {
		head.Id = 1
	}
	head.Data = data
}

func (head *DCircleLink) IsHeadEmpty() bool {
	return head.Next == nil && head.Prev == nil
}

func (head *DCircleLink) IsEmpty() bool {
	return head.Data == nil && head.Next == nil && head.Prev == nil
}

func (head *DCircleLink) GetLastNode() *DCircleLink {
	cur := head
	if !head.IsHeadEmpty() {
		for true {
			if cur.Next == head {
				break
			}
			cur = cur.Next
		}
	}
	return cur
}

func (head *DCircleLink) AddNode(newNode *DCircleLink) {
	if head.IsHeadEmpty() {
		head.Next = newNode
		head.Prev = newNode
		newNode.Prev = head
		newNode.Next = head
		return
	}
	cur := head
	flag := false // 标志，数据添加末尾
	for true {
		if cur == head.Prev { // 已经是最后一个节点，退出
			break
		} else if cur.Next.Id > newNode.Id { // 标志下数据应该插入到前列
			flag = true
			break
		} else if cur.Next.Id == newNode.Id { // 数据已经存在
			return
		}
		cur = cur.Next
	}
	if flag {
		// 最后一个节点，前面插入
		newNode.Next = cur.Next
		newNode.Prev = cur

		cur.Next.Prev = newNode
		cur.Next = newNode
	} else {
		newNode.Next = cur.Next
		newNode.Prev = cur

		cur.Next = newNode
		head.Prev = newNode
	}
}

func (head *DCircleLink) FindNodeById(id int) (*DCircleLink, bool) {
	if head.IsHeadEmpty() && head.Id == id {
		return head, true
	} else if head.IsHeadEmpty() && head.Id != id {
		return &DCircleLink{}, false
	}
	cur := head
	flag := false

	for true {
		if cur.Id == id {
			flag = true
			break
		}
		if cur == head.Prev {
			break
		}
		cur = cur.Next
	}
	if !flag {
		return &DCircleLink{}, false
	}
	return cur, true
}

func (head *DCircleLink) DeleteNodeById(id int) bool {
	if head.IsEmpty() {
		return false
	}
	node, ok := head.FindNodeById(id)
	if ok {
		// 删除头节点
		if node == head {
			// 只有一个节点
			if head.IsHeadEmpty() {
				head.Next = nil
				head.Prev = nil
				head.Data = nil
				head.Id = 0
				return true
			}
			// 只有2个节点
			if head.Next.Next == head {
				nextNode := head.Next

				head.Id = nextNode.Id
				head.Data = nextNode.Data
				head.Prev = nil
				head.Next = nil
				return true
			}
			// 移动下一个节点作为头节点
			nextNode := head.Next
			head.Data = nextNode.Data
			head.Id = nextNode.Id

			head.Next = nextNode.Next
			nextNode.Next.Prev = head
			return true
		}
		// 删除尾节点
		if node == head.GetLastNode() {
			// 只有2个节点
			if node.Prev == head && node.Next == head {
				head.Prev = nil
				head.Next = nil
				return true
			}
			head.Prev = node.Prev
			node.Prev.Next = head
			return true
		}
		// 删除中间节点
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		return true
	}
	return ok
}

func (head *DCircleLink) ChangeNodeById(id int, data interface{}) bool {
	node, ok := head.FindNodeById(id)
	if ok {
		node.Data = data
	}
	return ok
}

func (head *DCircleLink) Show() {
	if head.IsEmpty() {
		return
	}
	if head.IsHeadEmpty() {
		fmt.Println(head.Id, head.Data)
		return
	}
	cur := head
	for true {
		fmt.Print(cur.Id, cur.Data, "->")
		if cur == head.Prev {
			break
		}
		cur = cur.Next
	}
	fmt.Println("head")
}

func NewDCircleLinkNode(data interface{}) *DCircleLink {
	return &DCircleLink{1, data, nil, nil}
}
