package HashTableArray

import "errors"

const (
	Deleted = iota // 数据已被删除
	MinTableSize = 100 // 哈希表大小
	Legimate = iota // 已经存在的合法数据
	Empty = iota // 数据为空
)

// 哈希函数自定义
func SHA(str interface{}, tableSize int) int {
	var HashVar = 0
	var chars[] byte
	if strings, ok := str.(string); ok {
		chars = []byte(strings) // 字符串转化为字节数组
	}
	for _, v := range chars {
		HashVar = (HashVar << 17|123 & 1235^139) + int(v) // 哈希算法
	}
	return HashVar%MinTableSize
}

type HashFunc func(data interface{}, tableSize int) int // 函数指针
type HashEntry struct {
	data interface{}
	kind int // 类型
}

type HashTable struct {
	tableSize int // 哈希表的大小
	theCell []*HashEntry // 数组，每一个元素是指针指向哈希结构
	hashfunc HashFunc // 哈希函数
}

type HashTableImple interface {
	Find(data interface{}) int
	Insert(data interface{})
	Empty()
	GetValue(index int) interface{}
}

func NewHashTable(size int, hash HashFunc) (*HashTable, error) {
	if size < MinTableSize {
		return nil, errors.New("too less size")
	}
	if hash == nil {
		return nil, errors.New("hash function is nil")
	}
	hashtable := new(HashTable)
	hashtable.tableSize = size
	hashtable.theCell = make([]*HashEntry, size)
	hashtable.hashfunc = hash
	for i := 0; i < size; i++ {
		hashtable.theCell[i] = new(HashEntry)
		hashtable.theCell[i].data = nil
		hashtable.theCell[i].kind = Empty
	}
	return hashtable, nil
}

func (h HashTable) Find(data interface{}) int {
	collid := 0
	curpos := h.hashfunc(data, h.tableSize)
	if h.theCell[curpos].kind != Empty && h.theCell[curpos].data != data {
		collid += 1
		curpos = 2*curpos-1 // 平方探测, 冲突处理函数
		if curpos > h.tableSize {
			curpos -= h.tableSize // 越界处理
		}
	}
	return curpos
}

func (h HashTable) Insert(data interface{}) {
	pos := h.Find(data)
	entry := h.theCell[pos]
	if entry.kind != Legimate {
		entry.kind = Legimate
		entry.data = data
	}
}

func (h HashTable) Empty() {
	for i := 0; i < h.tableSize; i++ {
		if h.theCell[i] == nil {
			continue
		}
		h.theCell[i].kind = Deleted
	}
}

func (h HashTable) GetValue(index int) interface{} {
	if index > h.tableSize {
		return nil
	}
	entry := h.theCell[index]
	if entry.kind == Legimate {
		return entry.data
	}
	return nil
}

