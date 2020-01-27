package QueueLink

type LinKQueue interface {
	Length() int
	EnQueue(data interface{})
	DeQueue() interface{}
}

type Node struct {
	data interface{}
	pNext *Node
}

type QueueLink struct {
	front *Node
	rear *Node
}

func (q *QueueLink) Length() int {
	p := q.front
	length := 0
	for p.pNext != nil {
		p = p.pNext
		length++
	}
	return length
}

func (q *QueueLink) EnQueue(data interface{}) {
	NewNode := &Node{data, nil}
	if q.front == nil {
		q.front = NewNode
		q.rear = NewNode
	} else {
		q.rear.pNext = NewNode
		q.rear = q.rear.pNext
	}
}

func (q *QueueLink) DeQueue() interface{} {
	if q.front == nil {
		return nil
	}
	FrontNode := q.front
	if q.front == q.rear {
		q.front = nil
		q.rear = nil
	} else {
		q.front = q.front.pNext
	}
	return FrontNode.data
}



