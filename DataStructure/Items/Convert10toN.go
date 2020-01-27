package Items

import (
	"DataStructure/StackArray"
	"fmt"
)

// 10进制->n进制
func Convert10toN(x int, n int) {
	Stack := StackArray.NewStack()
	for x != 0 {
		Stack.Push(x%n)
		x /= 2
	}

	for Stack.Top() != nil {
		fmt.Print(Stack.Pop().(int))
	}
}
