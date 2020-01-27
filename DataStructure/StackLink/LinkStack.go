package StackLink

type Stack struct {
	data interface{}
	pNext *Stack
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewStack() *Stack {
	return new(Stack)
}

func (stack *Stack) IsEmpty() bool {
	if stack.pNext == nil {
		return true
	}
	return false
}

func (stack *Stack) Push(data interface{}) {
	NewNode := &Stack{data, nil}
	NewNode.pNext = stack.pNext
	stack.pNext = NewNode
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	data := stack.pNext.data
	stack.pNext = stack.pNext.pNext
	return data
}

func (stack *Stack) Length() int {
	p := stack
	length := 0
	for p.pNext != nil {
		p = p.pNext
		length++
	}
	return length
}





