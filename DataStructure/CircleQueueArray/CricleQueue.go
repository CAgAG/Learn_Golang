package CircleQueueArray

import "errors"

const QueueSize  = 5

type CQueue interface {
	QueueLength() int
	EnQueue(data interface{}) error
	DeQueue() (interface{}, error)
}


type CircleQueue struct {
	data [QueueSize] interface{}
	front int
	rear int
}

func NewCircleQueue(cq *CircleQueue) {
	cq.front = 0
	cq.rear = 0
}

func (cq *CircleQueue) QueueLength() int {
	return (cq.rear-cq.front + QueueSize) % QueueSize
}

func (cq *CircleQueue) EnQueue(data interface{}) error {
	if (cq.rear + 1)%QueueSize == cq.front % QueueSize{
		return errors.New("队列已满")
	}
	cq.data[cq.rear] = data
	cq.rear = (cq.rear + 1) % QueueSize
	return nil
}

func (cq *CircleQueue) DeQueue() (interface{}, error) {
	if  cq.rear==cq.front{
		return nil ,errors.New("队列为空")
	}
	res:=cq.data[cq.front] //取出第一个数据
	cq.data[cq.front]=0 //清空数据
	cq.front=(cq.front+1)% QueueSize
	return res,nil
}



