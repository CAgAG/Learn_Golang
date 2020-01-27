package ArrayList

import (
	"errors"
	"fmt"
)

// 接口
type List interface {
	Size() int // 长度
	Get(index int) (interface{}, error)// 得到值 interface{}为泛型
	Set(index int, newval interface{}) error
	Insert(index int, newval interface{}) error
	Append(newval interface{})
	Clear()
	Delete(index int)
	String() string
	ToDataStore()
	Iterator()
}

// 结构
type ArrayList struct {
	dataStore [] interface{} // 数值存储
	theSize int // 数值大小
}

func (List *ArrayList) Size() int {
	return List.theSize
}

func (List *ArrayList)Get(index int) (interface{}, error) {
	if index<0 || index>=List.theSize {
		return nil, errors.New("index out of Array")
	}
	return List.dataStore[index], nil
}

func (List *ArrayList) Set(index int, newval interface{}) error {
	if  index<0  || index >=List.theSize{
		return errors.New("索引越界")
	}
	List.dataStore[index] = newval
	return nil
}

func (List *ArrayList)checkFull() {
	if List.theSize == cap(List.dataStore) {
		newDataStore := make([]interface{}, 2*List.theSize, 2*List.theSize)
		copy(newDataStore, List.dataStore)
		List.dataStore = newDataStore
	}

}

func (List *ArrayList) Insert(index int, newval interface{}) error {
	if  index<0  || index >=List.theSize{
		return errors.New("索引越界")
	}
	List.checkFull()
	List.dataStore = List.dataStore[:List.theSize+1]
	for i:=List.theSize;i>index ;i-- {
		List.dataStore[i] = List.dataStore[i-1]
	}
	List.dataStore[index] = newval
	List.theSize++
	return nil
}

func (List *ArrayList) Append(newval interface{}) {
	List.dataStore = append(List.dataStore, newval)
	List.theSize++
}

func (List *ArrayList) Clear() {
	List.dataStore = make([]interface{}, 0, 10)
	List.theSize = 0
}

func (List *ArrayList) Delete(index int)  {
	List.dataStore = append(List.dataStore[:index], List.dataStore[index+1:]...)
	List.theSize--
}

func (List *ArrayList) String() string {
	return fmt.Sprint(List.dataStore)
}

func (List *ArrayList) ToDataStore() {
	fmt.Println(List.dataStore)
	fmt.Println(cap(List.dataStore))
}

func NewArrayList() *ArrayList  {
	List := new(ArrayList)
	List.dataStore = make([] interface{}, 0, 10)
	fmt.Println(List.dataStore)
	List.theSize = 0
	return List
}

func (list *ArrayList)Iterator() Iterator {
	it := new(ArrayListIterator)
	it.currentIndex = 0
	it.List = list
	return it
}
