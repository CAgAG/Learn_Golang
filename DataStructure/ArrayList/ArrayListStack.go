package ArrayList

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	array *ArrayList
	capsize int
}

func NewArrayListStack() *Stack {
	myStack := new(Stack)
	myStack.array = NewArrayList()
	myStack.capsize = 10
	return myStack
}

func (mystack *Stack) Clear() {
	mystack.array.Clear()
	mystack.capsize = 10
}

func (mystack *Stack) Size() int {
	return mystack.array.theSize
}

func (mystack *Stack) Pop() interface{} {
	if !mystack.IsEmpty() {
		last := mystack.array.dataStore[mystack.array.theSize-1]
		mystack.array.Delete(mystack.array.theSize-1)
		return last
	}
	return nil
}

func (mystack *Stack) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.array.Append(data)
	}
}

func (mystack *Stack) IsFull() bool {
	if mystack.array.theSize >= mystack.capsize {
		return true
	}
	return false
}

func (mystack *Stack) IsEmpty() bool {
	if mystack.array.theSize == 0 {
		return true
	}
	return false
}



