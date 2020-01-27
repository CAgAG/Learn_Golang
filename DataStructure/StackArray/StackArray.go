package StackArray

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
	Top() interface{}
}

type Stack struct {
	dataSource []interface{}
	capsize int // 最大范围
	cursize int // 当前大小
}


func NewStack() *Stack {
	myStack := new(Stack)
	myStack.dataSource = make([]interface{}, 0, 1000)
	myStack.capsize = 1000
	myStack.cursize = 0
	return myStack
}

func (MStack *Stack) Clear() {
	MStack.dataSource = make([]interface{}, 0, 1000)
	MStack.cursize = 0
	MStack.capsize = 1000
}

func (MStack *Stack) Pop() interface{} {
	if !MStack.IsEmpty() {
		last := MStack.dataSource[MStack.cursize-1]
		MStack.dataSource = MStack.dataSource[: MStack.cursize-1]
		MStack.cursize--
		return last
	}
	return nil
}

func (MStack *Stack) Push(data interface{}) {
	if !MStack.IsFull() {
		MStack.dataSource = append(MStack.dataSource, data)
		MStack.cursize++
	}
}

func (MStack *Stack) IsFull() bool {
	if MStack.cursize >= MStack.capsize {
		return true
	} else {
		return false
	}
}

func (MStack *Stack) IsEmpty() bool {
	if MStack.cursize == 0 {
		return true
	} else {
		return false
	}
}

func (MStack *Stack)Size() int {
	return MStack.cursize
}

func (MStack *Stack)Top() interface{} {
	if MStack.Size() == 0 {
		return nil
	}
	return MStack.dataSource[MStack.cursize-1]
}


