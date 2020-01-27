package SetLink

type Object interface {}
type MatchFunc func(data1, data2 Object) int // 对比函数

type Node struct {
	data Object
	next *Node
}

func (node *Node) GetData() Object {
	return node.data
}

type List struct {
	size uint64
	head *Node
	tail *Node
	Matcher MatchFunc
}

func (list *List) Match(data1 Object, data2 Object) int {
	var Matcher MatchFunc = nil
	if list.Matcher == nil {
		Matcher = DefaultMatch
	} else {
		Matcher = list.Matcher
	}
	return Matcher(data1, data2)
}

func (list *List) CreateNode(data Object) *Node {
	node := new(Node)
	node.data = data
	node.next = nil
	return node
}

func (list *List) NextNode(node *Node) *Node {
	return node.next
}

func (list *List) GetHead() *Node {
	return list.head
}

func (list *List) GetTail() *Node {
	return list.tail
}

func (list *List) RemoveAt(index uint64) Object {
	size := list.size
	if index >= size {
		return nil
	} else if size == 1 { // 仅有头节点
		node := list.head
		list.head = nil
		list.tail = nil
		list.size = 0
		return node.data
	} else if index == 0 { // 移除头节点
		node := list.head
		list.head = list.head.next

		list.size--
		return node.data
	} else if size == size-1 { // 移除尾节点
		pre := list.tail
		for i := uint64(2); i < size; i++ {
			pre = pre.next
		}
		tail := list.tail
		list.tail = pre
		pre.next = nil

		list.size--
		return tail.data
	} else {
		pre := list.head
		for i := uint64(2); i < index; i++ {
			pre = pre.next
		}
		node := pre.next
		Next := node.next

		node.next = Next
		list.size--
		return node.data
	}
}

func (list *List) Remove(data Object) bool {
	if data == nil || list.IsEmpty() {
		return false
	}
	head := list.head

	// 移除头节点
	if list.Matcher(head.GetData(), data) == 0 {
		list.head = list.head.next
	} else {
		pre := head
		cur := head.next
		for cur != nil {
			if list.Match(data, cur.GetData()) == 0 {
				pre.next = cur.next
				break
			}

			pre = cur
			cur = cur.next
		}
		if cur == nil {
			return false
		}
	}
	list.size--
	return true
}

func (list *List) IsMember(data Object) bool {
	if list.IsEmpty() {
		return false
	}
	head := list.head
	for i := head; i != nil; i = i.next {
		if list.Match(data, i.GetData()) == 0 {
			return true
		}
	}

	return false
}

func (list *List) Init(MatcherList ...MatchFunc) {
	list.size = 0
	list.head = nil
	list.tail = nil

	if len(MatcherList) == 0 {
		list.Matcher = DefaultMatch
	} else {
		list.Matcher = MatcherList[0]
	}
}

func (list *List) GetSize() uint64 {
	return list.size
}

func (list *List) IsEmpty() bool {
	return list.size == 0
}

func (list *List) Append(data Object) bool {
	newNode := new(Node)
	newNode.data = data
	newNode.next = nil

	if list.size == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		tail := list.tail
		tail.next = newNode
		list.tail = newNode
	}
	list.size++
	return true
}

func (list *List) InsertAtHead(data Object) bool {
	newNode := list.CreateNode(data)

	newNode.next = list.head
	list.head = newNode

	list.size++
	return true
}

func (list *List) First() Object {
	if list.size == 0 {
		return nil
	} else {
		return list.head.data
	}
}

func (list *List) Last() Object {
	if list.size == 0 {
		return nil
	} else {
		return list.tail.data
	}
}

func (list *List) Next(curData Object) Object {
	head := list.head
	for i := head; i != nil; i = i.next {
		if list.Match(curData, i.GetData()) == 0 {
			Next := i.next
			if Next == nil {
				return nil
			} else {
				return Next.GetData()
			}
		}
	}
	return nil
}

func (list *List) GetAt(index uint64) Object {
	size := list.GetSize()

	if index >= size || index < 0 {
		return nil
	} else if index == 0 {
		return list.First()
	} else if index == size-1 {
		return list.Last()
	} else {
		cur := list.head
		for i := uint64(0); i < size; i++ {
			if i == index {
				break
			}
			cur = cur.next
		}
		return cur.GetData()
	}
}

func (list *List) InsertAt(index uint64, data Object) bool {
	size := list.GetSize()
	if index > size || index < 0 {
		return false
	} else if index == size {
		return list.Append(data)
	} else if index == 0 {
		return list.InsertAtHead(data)
	} else {
		newNode := list.CreateNode(data)
		preI := index - 1
		cur := list.head
		for i := uint64(0); i < size; i++ {
			if i == preI {
				break
			}
			cur = cur.next
		}

		newNode.next = cur.next
		cur.next = newNode

		list.size++
		return true
	}
}

func (list *List) Clear() {
	list.Init()
}

type ListLinkImple interface {
	Match(data1 Object, data2 Object) int
	CreateNode(data Object) *Node
	NextNode(node *Node) *Node
	GetHead() *Node
	GetTail() *Node
	RemoveAt(index uint64) Object
	Remove(data Object) bool
	IsMember(data Object) bool
	Init(yourMatch ...MatchFunc)
	GetSize() uint64
	IsEmpty() bool
	Append(data Object) bool
	InsertAtHead(data Object) bool
	First() Object
	Last() Object
	Next(curData Object) Object
	GetAt(index uint64) Object
	InsertAt(index uint64, data Object) bool
	Clear()
}

func DefaultMatch(data1, data2 Object) int {
	if data2 == data1 {
		return 0
	}
	return 1
}
















