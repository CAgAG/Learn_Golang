package ArrayList

type Iterator interface {
	HasNext() bool
	Next()(interface{}, error)
	Remove()
	GetIndex() int
}

//type Iterable interface {
//	Iterator() Iterator
//}

type ArrayListIterator struct {
	List *ArrayList
	currentIndex int
}

func (it *ArrayListIterator)HasNext() bool {
	return it.currentIndex < it.List.theSize
}

func (it *ArrayListIterator)Next() (interface{}, error) {
	if !it.HasNext(){
		return nil, nil
	}
	value, err := it.List.Get(it.currentIndex)
	it.currentIndex++
	return value, err

}

func (it *ArrayListIterator)Remove()  {
	it.currentIndex--
	it.List.Delete(it.currentIndex)
}

func (it *ArrayListIterator)GetIndex() int {
	return it.currentIndex
}
