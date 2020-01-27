package SetLink

type Set struct {
	LinkList *List
}

func (set *Set) GetAt(index uint64) Object {
	return set.LinkList.GetAt(index)
}

func (set *Set) GetSize() uint64 {
	return set.LinkList.GetSize()
}

func (set *Set) Init(matchers ...MatchFunc) {
	LinkList := new(List)
	set.LinkList = LinkList

	set.LinkList.Init(matchers[0])
}

func (set *Set) IsExist(data Object) bool {
	return set.LinkList.IsMember(data)
}

func (set *Set) Insert(data Object) bool {
	if !set.IsExist(data) {
		return set.LinkList.Append(data)
	}
	return false
}

func (set *Set) IsEmpty() bool {
	return set.LinkList.IsEmpty()
}

func (set *Set) Remove(data Object) bool {
	return set.LinkList.Remove(data)
}

// 并集
func (set *Set) Union(setT *Set) *Set {
	if setT == nil {
		return set
	}
	if set == nil {
		return setT
	}
	newSet := new(Set)
	newSet.Init(set.LinkList.Matcher)

	if set.IsEmpty() && setT.IsEmpty() {
		return newSet // 空集
	}
	// 复制set
	for i := uint64(0); i < set.GetSize(); i++ {
		newSet.Insert(set.GetAt(i))
	}
	var data Object
	for i := uint64(0); i < setT.GetSize(); i++ {
		data = setT.GetAt(i)
		if !newSet.IsExist(data) {
			newSet.Insert(data)
		}
	}
	return newSet
}

// 交集
func (set *Set) Share(setT *Set) *Set {
	if setT == nil {
		return set
	}
	if set == nil {
		return setT
	}
	newSet := new(Set)
	newSet.Init(set.LinkList.Matcher)

	if set.IsEmpty() && setT.IsEmpty() {
		return newSet // 空集
	}

	large := set // 保存最多元素
	small := setT  // 较少元素
	if setT.GetSize() > set.GetSize() {
		large = setT
		small = set
	}

	var data Object
	for i := uint64(0); i < large.GetSize(); i++ {
		data = large.GetAt(i)

		if small.IsExist(data) {
			newSet.Insert(data)
		}
	}
	return newSet
}

// 减运算
func (set *Set) Different(setT *Set) *Set {
	if setT == nil {
		return set
	}
	if set == nil {
		return setT
	}
	newSet := new(Set)
	newSet.Init(set.LinkList.Matcher)

	if set.IsEmpty() && setT.IsEmpty() {
		return newSet // 空集
	}

	var data Object
	for i := uint64(0); i < set.GetSize(); i++ {
		data = set.GetAt(i)

		if !setT.IsExist(data) {
			newSet.Insert(data)
		}
	}
	return newSet
}

func (set *Set) IsSub(subset *Set) bool {
	if set == nil {
		return false
	}
	if subset == nil {
		return true
	}
	for i := uint64(0);i < subset.GetSize(); i++ {
		if !set.IsExist(subset.GetAt(i)) {
			return false
		}
	}
	return true
}

func (set *Set) IsEquals(subset *Set) bool {
	if set == nil && subset == nil {
		return true
	}
	if set != nil && subset == nil {
		return false
	}
	if set == nil && subset != nil {
		return false
	}
	newSet := set.Share(subset)
	return newSet.GetSize() == set.GetSize()
}

func (set *Set) GetIterator() *SetIterator {
	it := new(SetIterator)
	it.index = 0
	it.set = set
	return it
}

type SetLinkImple interface {
	GetAt(index uint64) Object
	GetSize() uint64
	Init(match...MatchFunc)
	IsExist(data Object) bool
	Insert(data Object) bool
	IsEmpty() bool
	Remove(data Object) bool
	Union(setT *Set) *Set
	Share(setT *Set) *Set
	Different(setT *Set) *Set
	IsSub(subset *Set) bool
	IsEquals(subset *Set) bool
	GetIterator() *SetIterator
}
