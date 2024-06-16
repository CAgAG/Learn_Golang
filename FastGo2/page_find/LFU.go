package main

import (
	"container/list"
	"fmt"
)

type FreqItem struct {
	key   int
	value int
	count int
}

type LFUCache struct {
	capacity int
	cache    map[int]*list.Element
	freqList *list.List
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		freqList: list.New(),
	}
}

func (lfu *LFUCache) Get(key int) int {
	if node, ok := lfu.cache[key]; ok {
		node.Value.(*FreqItem).count++ // 访问次数 加一
		lfu.updateFreqList(node)       // 保持顺序
		return node.Value.(*FreqItem).value
	}
	return -1
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}
	// 如果已经存在
	if node, ok := lfu.cache[key]; ok {
		node.Value.(*FreqItem).value = value
		node.Value.(*FreqItem).count++
		lfu.updateFreqList(node)
	} else {
		if len(lfu.cache) >= lfu.capacity {
			evict := lfu.freqList.Back()
			delete(lfu.cache, evict.Value.(*FreqItem).key)
			lfu.freqList.Remove(evict)
		}
		newNode := &FreqItem{key: key, value: value, count: 1}
		element := lfu.freqList.PushBack(newNode)
		lfu.cache[key] = element
	}
}

func (lfu *LFUCache) updateFreqList(node *list.Element) {
	for e := lfu.freqList.Front(); e != nil; e = e.Next() {
		if e.Value.(*FreqItem).count < node.Value.(*FreqItem).count { // 从 大到小 排序
			lfu.freqList.MoveBefore(node, e) // 将 node 移到 e 的前面
			return
		}
	}
}

func main() {
	cache := NewLFUCache(3)
	cache.Put(1, 1)
	cache.Put(2, 2)

	for e := cache.freqList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*FreqItem).value)
	}
	fmt.Println("=========================================")

	cache.Put(3, 3)
	for e := cache.freqList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*FreqItem).value)
	}
	fmt.Println("=========================================")

	cache.Put(4, 4)
	for e := cache.freqList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*FreqItem).value)
	}
	fmt.Println("=========================================")

	fmt.Println("Get: ", cache.Get(4))
	for e := cache.freqList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*FreqItem).value)
	}
	fmt.Println("=========================================")
	cache.Put(2, 22)
	for e := cache.freqList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*FreqItem).value)
	}
	fmt.Println("=========================================")
	cache.Put(5, 5)
	for e := cache.freqList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*FreqItem).value)
	}
	fmt.Println("=========================================")
}
