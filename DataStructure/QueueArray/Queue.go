package QueueArray

type MQueue interface {
	Size() int
	Front() interface{}
	Tail() interface{}
	IsEmpty() bool
	EnQueue(data interface{})
	DeQueue() interface{}
	Clear()
	Shift() interface{}
}

type Queue struct {
	dataSource []interface{}
	thsSize int
}

func (q *Queue) Size() int {
	return q.thsSize
}

func (q *Queue) Front() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.dataSource[0]
}

func (q *Queue) Tail() interface{} {
	if q.thsSize == 0 {
		return nil
	}
	return q.dataSource[q.Size()-1]
}

func (q *Queue) IsEmpty() bool {
	return q.thsSize == 0
}

func (q *Queue) EnQueue(data interface{}) {
	q.dataSource = append(q.dataSource, data)
	q.thsSize++
}

func (q *Queue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := q.dataSource[0]
	if q.Size() > 1 {
		q.dataSource = q.dataSource[1: q.Size()]
	}else {
		q.dataSource = make([]interface{}, 0)
	}
	q.thsSize--
	return data
}

func (q *Queue) Clear() {
	q.dataSource = make([]interface{}, 0)
	q.thsSize = 0
}

func (q *Queue)Shift() interface{} {
	data := q.dataSource[0]
	q.dataSource = q.dataSource[1: ]
	q.thsSize--
	return data
}

func NewQueue() *Queue {
	myqueue := new(Queue)
	myqueue.dataSource = make([]interface{}, 0)
	myqueue.thsSize = 0
	return myqueue
}


