package HashTableLink

import "errors"

type List struct {
	length int
	Head *Node
	Tail *Node
}

func (list *List) Len() int {
	return list.length
}

func (list *List) IsEmpty() bool {
	return list.length == 0
}

func (list *List) Prepend(value interface{}) {
	node := NewNode(value)
	if list.length == 0 {
		list.Head = node
		list.Tail = node
	} else {
		FHead := node
		FHead.Pre = node

		node.Next = FHead
		list.Head = node
	}
}

func (list *List) Append(value interface{}) {
	node := NewNode(value)
	if list.length == 0 {
		list.Head = node
		list.Tail = node
	} else {
		FTail := node
		FTail.Next = node

		node.Pre = FTail
		list.Tail = node
	}
	list.length++
}

func (list *List) Add(value interface{}, index int) error {
	if index > list.length {
		return errors.New(" index out of range")
	}
	node := NewNode(value)
	if list.length == 0 || index == 0{
		list.Prepend(value)
		return nil
	}
	if list.length-1 == index {
		list.Append(value)
		return nil
	}
	NextNode, _ := list.Get(index)
	PreNode := NextNode.Pre

	PreNode.Next = node
	node.Pre = PreNode

	NextNode.Pre = node
	node.Next = NextNode

	list.length++
	return nil
}

func (list *List) Remove(value interface{}) error {
	if list.length == 0 {
		return errors.New("empty list")
	}
	if list.Head.Value == value {
		list.Head = list.Head.Next
		list.length--
		return nil
	}
	flag := true
	for n := list.Head; n != nil; n = n.Next {
		if n.Value == value && flag {
			n.Next.Pre, n.Pre.Next = n.Pre, n.Next
			list.length--
			flag = false
		}
	}
	if !flag {
		return errors.New("not found")
	}
	return nil
}

func (list *List) Get(index int) (*Node, error) {
	if index > list.length {
		return nil, errors.New("index out of range")
	}
	node := list.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node, nil
}

func (list *List) Find(node *Node) (int, error) {
	if list.length == 0 {
		return 0, errors.New("empty list")
	}
	index := 0
	found := -1
	list.Map(func(n *Node) {
		index++
		if n.Value == node.Value && found == -1 {
			found = index
		}
	})
	if found == -1 {
		return 0, errors.New("not found")
	}
	return found, nil
}

func (list *List) Clear() {
	list.Head = nil
	list.Tail = nil
	list.length = 0
}

func (list *List) Concat(k *List) {
	list.Tail.Next, k.Head.Pre = k.Head, list.Tail
	list.Tail = k.Tail
	list.length += k.length
}

func (list *List) Map(f func(node *Node)) {
	for node := list.Head; node != nil; node=node.Next {
		n := node.Value.(*Node)
		f(n)
	}
}

func (list *List) Each(f func(node Node)) {
	for node := list.Head; node != nil; node=node.Next {
		f(*node)
	}
}

type Node struct {
	Value interface{}
	Pre *Node
	Next *Node
}

type ListImple interface {
	Len() int
	IsEmpty() bool
	Prepend(value interface{})
	Append(value interface{})
	Add(value interface{}, index int) error
	Remove(value interface{}) error
	Get(index int) (*Node, error)
	Find(node *Node) (int, error)
	Clear()
	Concat(k *List)
	Map(f func(node *Node))
	Each(f func(node Node))
}

func NewNode(value interface{}) *Node {
	return &Node{value, nil, nil}
}

func NewList() *List {
	list := new(List)
	list.length = 0
	return list
}


