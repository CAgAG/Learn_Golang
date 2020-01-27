package HeapArray

import "sync"

type DataType int
type Item interface {
	Less(than Item) bool
}

func (x DataType)Less(than Item) bool {
	return x < than.(DataType)
}

type Heap struct {
	lock *sync.Mutex // 线程安全
	data []Item
	min bool // 是否小顶堆
}

type HeapImple interface {
	IsEmpty() bool
	Len() int
	Get(index int) Item
	Insert(it Item)
	Less(a, b Item) bool
	Extract() Item
	SiftUp()
	SiftDown()
}

// 小顶堆
func NewHeap() *Heap {
	return &Heap{new(sync.Mutex), make([]Item, 0), true}
}

// 大顶堆
func NewMaxHeap() *Heap {
	return &Heap{new(sync.Mutex), make([]Item, 0), false}
}


func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Get(index int) Item {
	return h.data[index]
}

func (h *Heap) Insert(it Item) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.data = append(h.data, it)
	h.SiftUp()
	return
}

func (h *Heap) Less(a, b Item) bool {
	if h.min {
		return a.Less(b)
	} else {
		return b.Less(a)
	}
}

// 压缩，弹出一个
func (h *Heap) Extract() Item {
	h.lock.Lock()
	defer h.lock.Unlock()

	if h.Len() == 0 {
		return nil
	}
	length := h.Len()
	first := h.data[0]
	last := h.data[length-1]
	if length == 1 {
		h.data = nil // 重新分配内存
		return nil
	}
	h.data = append([]Item{last}, h.data[1: length-1]...)
	h.SiftDown()
	return first
}

// 弹出一个极大值
func (h *Heap) SiftUp() {
	// 堆排序的循环过程  n,2n+1
	length := h.Len()
	for i, parent := length-1, length-1; i > 0; i=parent {
		parent = i/2
		if h.Less(h.Get(i), h.Get(parent)) {
			h.data[parent], h.data[i] = h.data[i], h.data[parent]
		} else {
			break
		}
	}
}

// 弹出一个极小值
func (h *Heap) SiftDown() {
	// 堆排序的循环过程  n,2n+1
	length := h.Len()
	for i, child := 0, 1; i < length && 2*i+1 < length; i = child {
		child = 2*i+1
		if child+1 <= length-1 && h.Less(h.Get(child+1), h.Get(child)) {
			child++  // 循环左右节点过程
		}
		if h.Less(h.Get(i), h.Get(child)) {
			break
		}
		h.data[i],h.data[child]=h.data[child],h.data[i]
	}
}






















