package DeQueArray

import (
	"fmt"
	"sync"
)

type Deque struct {
	arr         []interface{}
	left, right int
	fixSize     int
	lock        *sync.RWMutex
}

type DequeImple interface {
	Addleft(data interface{})
	Addright(data interface{})
	Delleft() interface{}
	Delright() interface{}
	Show()
}

func (deQ *Deque) Addleft(data interface{}) {
	if deQ.right == deQ.left && deQ.left != 0 {
		panic("overflow")
	}
	deQ.lock.Lock()
	defer deQ.lock.Unlock()

	deQ.arr[deQ.left] = data
	deQ.left -= 1
	if deQ.left == -1 {
		deQ.left = deQ.fixSize - 1 // 循环双端队列
	}
}

func (deQ *Deque) Addright(data interface{}) {
	if deQ.right == deQ.left && deQ.right != 0 {
		panic("overflow")
	}
	deQ.lock.Lock()
	defer deQ.lock.Unlock()

	deQ.arr[deQ.right] = data
	deQ.right += 1
	if deQ.right == deQ.fixSize {
		deQ.right = 0
	}
}

func (deQ *Deque) Delleft() interface{} {
	if deQ.fixSize == deQ.left {
		panic("overflow")
	}
	deQ.lock.Lock()
	defer deQ.lock.Unlock()

	deQ.left += 1
	if deQ.left == deQ.fixSize { // 到尾部
		deQ.left = 0
	}
	data := deQ.arr[deQ.left]
	return data
}

func (deQ *Deque) Delright() interface{} {
	if deQ.right == deQ.left {
		panic("overflow")
	}
	deQ.lock.Lock()
	defer deQ.lock.Unlock()

	deQ.right -= 1
	if deQ.right == -1 {
		deQ.right = deQ.fixSize - 1
	}
	data := deQ.arr[deQ.right]
	return data
}

func (deQ *Deque) Show() {
	fmt.Println(deQ.arr)
	fmt.Println("left:", deQ.left, "right:", deQ.right)
}

func NewDeque(cap int) *Deque {
	if cap <= 0 {
		panic("must be greater than 0")
	}
	deQ := new(Deque)
	deQ.arr = make([]interface{}, cap)
	deQ.fixSize = cap
	deQ.left = 0
	deQ.right = 0
	deQ.lock = new(sync.RWMutex)
	return deQ
}
