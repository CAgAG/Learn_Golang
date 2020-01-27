package DoubleLink

type DoubleLinkNode struct {
	value interface{}
	pre *DoubleLinkNode
	next *DoubleLinkNode
}

func NewDoubleLinkNode(value interface{}) *DoubleLinkNode {
	return &DoubleLinkNode{value, nil, nil}
}

func (node *DoubleLinkNode)Value() interface{} {
	return node.value
}

func (node *DoubleLinkNode)Pre() interface{} {
	return node.pre
}

func (node *DoubleLinkNode)Next() interface{} {
	return node.next
}
