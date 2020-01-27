package QueueArray

import "DataStructure/HeapArray"

type PriorityItem struct {
	value interface{}
	Priority int
}

func NewPriorityitem(value interface{}, Priority int) *PriorityItem {
	return &PriorityItem{value, Priority}
}

func (x PriorityItem)Less(than HeapArray.Item) bool {
	return x.Priority<than.(PriorityItem).Priority
}

// 优先队列，基于堆
type PriorityQueue struct {
	data *HeapArray.Heap
}

type PriorityQueueImple interface {
	Len() int
	Insert(data PriorityItem)
	Extract() PriorityItem
	ChangePriority(val interface{}, Priority int)
}

func NewMaxPriorityQueue() *PriorityQueue {
	return &PriorityQueue{HeapArray.NewMaxHeap()}
}

func NewMinPriorityQueue() *PriorityQueue {
	return &PriorityQueue{HeapArray.NewHeap()}
}

func (p PriorityQueue) Len() int {
	return p.data.Len()
}

func (p PriorityQueue) Insert(data PriorityItem) {
	p.data.Insert(data)
}

func (p PriorityQueue) Extract() PriorityItem {
	return p.data.Extract().(PriorityItem)
}

func (p PriorityQueue) ChangePriority(val interface{}, Priority int) {
	storage := NewQueue()
	popped := p.Extract() // 拿出最小值
	for val != popped.value {
		if p.Len() == 0 {
			return
		}
		storage.EnQueue(popped)
		popped = p.Extract()
	}
	popped.Priority = Priority
	p.data.Insert(popped)
	for storage.Size() > 0 {
		p.data.Insert(storage.Shift().(HeapArray.Item))
	}
}




