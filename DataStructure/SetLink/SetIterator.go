package SetLink

type SetIterator struct {
	index uint64
	set *Set
}

func (it *SetIterator) HashNext() bool {
	setI := it.set
	index := it.index
	return index < setI.GetSize()
}

func (it *SetIterator) Next() Object {
	setI := it.set
	index := it.index
	if index < setI.GetSize() {
		data := setI.GetAt(index)
		it.index++
		return data
	}
	return nil
}

type SetIteratorImple interface {
	HashNext() bool
	Next() Object
}
