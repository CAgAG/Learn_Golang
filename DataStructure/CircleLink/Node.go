package CircleLink

type Node struct {
	Data interface{}
	Next *Node
}

var Head, Tail *Node
