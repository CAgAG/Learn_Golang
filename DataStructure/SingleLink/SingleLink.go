package SingleLink

import (
	"fmt"
	"strings"
)

type SingleLink interface {
	GetFirstNode() *SingleLinkNode
	GetHead() *SingleLinkNode
	InsertNodeFront(node *SingleLinkNode)
	InsertNodeTail(node *SingleLinkNode)

	InsertNodeValueBack(dest interface{},node *SingleLinkNode ) bool
	InsertNodeValueFront(dest interface{},node *SingleLinkNode )bool

	GetNodeAtIndex(index int) *SingleLinkNode
	DeleteNode(node *SingleLinkNode) bool
	DeleteIndex(index int)

	String() string
}

type SingleLinkList struct {
	head *SingleLinkNode
	length int
}

func (list *SingleLinkList) GetFirstNode() *SingleLinkNode {
	return list.head.Next
}

func (list *SingleLinkList) GetHead() *SingleLinkNode {
	return list.head
}

func (list *SingleLinkList) InsertNodeFront(node *SingleLinkNode) {
	if list.head == nil {
		list.head.Next = node
		node.Next = nil
	} else {
		bak := list.head
		node.Next = bak.Next
		bak.Next = node
	}
	list.length++
}

func (list *SingleLinkList) InsertNodeTail(node *SingleLinkNode) {
	if list.head == nil {
		list.head.Next = node
		node.Next = nil
	} else {
		bak := list.head
		for bak.Next != nil {
			bak = bak.Next
		}
		bak.Next = node
	}
	list.length++
}

func (list *SingleLinkList) InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool {
	p := list.head
	isFind := false
	for p.Next != nil {
		if p.Value == dest {
			isFind = true
			break
		}
		p = p.Next
	}
	if isFind {
		node.Next = p.Next
		p.Next = node
		list.length++
	}
	return isFind
}

func (list *SingleLinkList) InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool {
	p := list.head
	isFind := false
	for p.Next != nil {
		if p.Next.Value == dest {
			isFind = true
			break
		}
		p = p.Next
	}
	if isFind {
		node.Next = p.Next
		p.Next = node
		list.length++
	}
	return isFind
}

func (list *SingleLinkList) GetNodeAtIndex(index int) *SingleLinkNode {
	if index > list.length-1 || index < 0 {
		return nil
	} else {
		p := list.head
		for index > -1 {
			p = p.Next
			index--
		}
		return p
	}
}

func (list *SingleLinkList) DeleteNode(node *SingleLinkNode) bool {
	if node == nil {
		 return false
	}
	p := list.head
	for p.Next != nil && p.Next != node {
		p = p.Next
	}
	if p.Next == node {
		p.Next = p.Next.Next
		list.length--
		return true
	}
	return false
}

func (list *SingleLinkList) DeleteIndex(index int) {
	if index > list.length-1 || index < 0 {
		return
	} else {
		p := list.head
		for index > 0 {
			p = p.Next
			index--
		}
		p.Next = p.Next.Next
		list.length--
		return
	}
}

func (list *SingleLinkList) String() string {
	var listString string
	p := list.head
	for p.Next != nil {
		listString += fmt.Sprintf("%v-->", p.Next.Value)
		p = p.Next
	}
	listString += fmt.Sprintf("nil")
	return listString
}

func NewSingleLinkList() *SingleLinkList {
	head := NewSingleLinkNode(nil)
	return &SingleLinkList{head, 0}
}

func (list *SingleLinkList)FindString(data string){
	p := list.head.Next
	for p.Next != nil {
		if strings.Contains(p.Value.(string), data) {
			fmt.Println(p.Value)
		}
		p = p.Next
	}
}

// 得到中间节点
func (list *SingleLinkList)GetMid() *SingleLinkNode{
	if list.head.Next == nil {
		return nil
	} else {
		p1 := list.head
		p2 := list.head
		for p2 != nil && p2.Next != nil {
			p1 = p1.Next
			p2 = p2.Next.Next
		}
		return p1
	}
}

func (list *SingleLinkList)ReverseList() {
	if list.head == nil || list.head.Next == nil {
		return
	} else {
		var (
			pre *SingleLinkNode
			cur = list.head.Next
		)
		for cur != nil {
			curNext := cur.Next
			cur.Next = pre

			pre = cur
			cur = curNext
		}
		list.head.Next.Next = nil
		list.head.Next = pre
	}
}


