package DoubleLink

import (
	"fmt"
	"strings"
)

type DoubleLinkList struct {
	head *DoubleLinkNode
	length int
}

type DoubleLink interface {
	GetLength() int
	GetFirstNode() *DoubleLinkNode
	InsertHead(node *DoubleLinkNode)
	InsertBack(node *DoubleLinkNode)

	InsertValueBack(dest  *DoubleLinkNode,node *DoubleLinkNode ) bool
	InsertValueHead(dest  *DoubleLinkNode,node *DoubleLinkNode ) bool
	InsertValueBackByValue(dest  interface{},node *DoubleLinkNode ) bool
	InsertValueHeadByValue(dest  interface{},node *DoubleLinkNode ) bool

	GetNodeAtIndex(index int) *DoubleLinkNode

	DeleteNode(node *DoubleLinkNode) bool
	DeleteNodeAtIndex(index int) bool

	FindString(inputstr string)
	String() string
}

func NewDoubleLinkList() *DoubleLinkList {
	head := NewDoubleLinkNode(nil)
	return &DoubleLinkList{head, 0}
}

func (dlist *DoubleLinkList) GetLength() int {
	return dlist.length
}

func (dlist *DoubleLinkList) GetFirstNode() *DoubleLinkNode {
	return dlist.head.next
}

func (dlist *DoubleLinkList) InsertHead(node *DoubleLinkNode) {
	p := dlist.head
	if p.next == nil {
		p.next = node
		node.pre = p
	} else {
		p.next.pre = node // 标记上一个节点
		node.next = p.next

		p.next = node
		node.pre = p
	}
	dlist.length++
}

func (dlist *DoubleLinkList) InsertBack(node *DoubleLinkNode) {
	p := dlist.head
	if p.next == nil {
		node.next = nil
		p.next = node
		node.pre = p
	} else {
		for p.next != nil {
			p = p.next
		}
		p.next = node
		node.pre = p
	}
	dlist.length++
}

func (dlist *DoubleLinkList) InsertValueBack(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	p := dlist.head
	for p.next != nil && p.next != dest {
		p = p.next
	}
	if p.next == dest {
		if p.next.next != nil {
			p.next.next.pre = node
		}
		node.next = p.next.next
		p.next.next = node
		node.pre = p.next
		dlist.length++
		return true
	}
	return false
}

func (dlist *DoubleLinkList) InsertValueHead(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	p := dlist.head
	for p.next != nil && p.next.next != dest {
		p = p.next
	}
	if p.next == dest {
		if p.next != nil {
			p.next.pre = node
		}
		node.next = p.next
		node.pre = p
		p.next = node

		dlist.length++
		return true
	}
	return false
}

func (dlist *DoubleLinkList) InsertValueBackByValue(dest interface{}, node *DoubleLinkNode) bool {
	p := dlist.head
	for p.next != nil && p.next.value != dest {
		p = p.next
	}
	if p.next.value == dest {
		if p.next.next != nil {
			p.next.next.pre = node
		}
		node.next = p.next.next
		p.next.next = node
		node.pre = p.next

		dlist.length++
		return true
	}
	return false
}

func (dlist *DoubleLinkList) InsertValueHeadByValue(dest interface{}, node *DoubleLinkNode) bool {
	p := dlist.head
	for p.next != nil && p.next.value != dest {
		p = p.next
	}
	if p.next.value == dest {
		if p.next != nil {
			p.next.pre = node
		}
		node.next = p.next
		node.pre = p
		p.next = node

		dlist.length++
		return true
	}
	return false
}

func (dlist *DoubleLinkList) GetNodeAtIndex(index int) *DoubleLinkNode {
	if index > dlist.length-1 || index < 0 {
		return nil
	}
	p := dlist.head
	for index > -1 {
		p = p.next
		index--
	}
	return p
}

func (dlist *DoubleLinkList) DeleteNode(node *DoubleLinkNode) bool {
	if node == nil {
		return false
	} else {
		p := dlist.head
		for p.next != nil && p.next != node {
			p = p.next
		}
		if p.next == node {
			if p.next.next != nil {
				p.next.next.pre = p
			}
			p.next = p.next.next

			dlist.length--
			return true
		}
		return false
	}
}

func (dlist *DoubleLinkList) DeleteNodeAtIndex(index int) bool {
	if index > dlist.length-1 || index < 0 {
		return false
	}
	p := dlist.head
	for index > 0 {
		p = p.next
		index--
	}
	if p.next.next != nil {
		p.next.next.pre = p
	}
	p.next = p.next.next
	dlist.length--
	return true
}

func (dlist *DoubleLinkList) FindString(inputstr string) {
	p := dlist.head.next
	for p.next != nil {
		if strings.Contains(p.value.(string), inputstr) {
			fmt.Println(p.value)
		}
		p = p.next
	}
}

func (dlist *DoubleLinkList) String() string {
	var (
		listString1 string
		listString2 string
	)
	p := dlist.head
	for p.next != nil {
		listString1 += fmt.Sprintf("%v-->", p.next.value)
		p = p.next
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n"
	for p != dlist.head {
		listString2 += fmt.Sprintf("<--%v", p.value)
		p = p.pre
	}
	listString1 += fmt.Sprintf("nil")
	return listString1+listString2+"\n"
}


