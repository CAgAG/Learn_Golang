package SingleLink

import "fmt"

type SingleLinkNode struct {
	Value interface{}
	Next  *SingleLinkNode
}

func NewSingleLinkNode(data interface{}) *SingleLinkNode {
	return &SingleLinkNode{data, nil}
}

func (node *SingleLinkNode) SwapValue (Node *SingleLinkNode) {
	node.Value, Node.Value = Node.Value, node.Value
}

func (head *SingleLinkNode) PrintLinkByNode() string {
	p := head.Next
	s := ""
	for p != nil {
		s += fmt.Sprintf("%v-->", p.Value)
		p = p.Next
	}
	s += "nil"
	return s
}
