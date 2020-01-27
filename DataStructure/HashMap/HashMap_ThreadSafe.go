package HashMap

import (
	"hash/crc32"
	"sort"
	"sync"
)

type Rindex []uint32 // hash 环索引

type Ring struct {
	Rmap      map[uint32]string
	RindexArr Rindex
	Lock      *sync.RWMutex
}

type HashMapImple interface {
	AddNode(nodeName string)
	RemoveNode(nodeName string)
	GetNode(nodeName string) string
}

// sort imple
// 比较大小
func (this Rindex) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Rindex) Len() int {
	return len(this)
}

func (this Rindex) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
// -------------------------------------

func (this *Ring) AddNode(nodeName string) {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	index := crc32.ChecksumIEEE([]byte(nodeName)) // sha256
	// 有值 退出
	if _, ok := this.Rmap[index]; ok {
		return
	}
	// 空
	this.Rmap[index] = nodeName                    // 赋值
	this.RindexArr = append(this.RindexArr, index) // 加载索引
	sort.Sort(this.RindexArr)                      // 排序
}

func (this *Ring) RemoveNode(nodeName string) {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	index := crc32.ChecksumIEEE([]byte(nodeName)) // sha256
	if _, ok := this.Rmap[index]; !ok {
		return
	}

	delete(this.Rmap, index)
	// 重新加载索引
	this.RindexArr = Rindex{}
	for k := range this.Rmap {
		this.RindexArr = append(this.RindexArr, k)
	}
	sort.Sort(this.RindexArr)
}

func (this *Ring) GetNode(nodeName string) string {
	this.Lock.RLock() //其他线程可以读取，不可以修改
	defer  this.Lock.RUnlock()

	hash := crc32.ChecksumIEEE([]byte(nodeName))

	index := -1
	for i, v := range this.RindexArr {
		if v == hash {
			index = i
			break
		}
	}

	if index == -1 {
		return ""
	}
	node := this.Rmap[this.RindexArr[index]] // 取得节点
	return node
}
