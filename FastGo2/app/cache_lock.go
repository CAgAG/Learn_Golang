package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// CacheEntry 表示缓存中的一个条目
type CacheEntry struct {
	key        string
	value      interface{}
	lastUsed   time.Time
	expiration time.Duration
}

// LRUCache 带有超时功能的LRU缓存
type LRUCache struct {
	cap   int
	data  *list.List
	cache map[string]*list.Element
	lock  sync.Mutex
}

// NewLRUCache 创建一个新的LRU缓存
func NewLRUCache(capacity int) *LRUCache {
	l := list.New()
	return &LRUCache{
		cap:   capacity,
		data:  l,
		cache: make(map[string]*list.Element),
	}
}

// Get 从缓存中获取一个值
func (lru *LRUCache) Get(key string) (interface{}, bool) {
	lru.lock.Lock()
	defer lru.lock.Unlock()

	element, ok := lru.cache[key]
	if !ok {
		return nil, false
	}

	entry := element.Value.(*CacheEntry)
	if time.Now().After(entry.lastUsed.Add(entry.expiration)) {
		// 条目已过期，从缓存中移除
		lru.removeElement(element)
		return nil, false
	}
	entry.lastUsed = time.Now()
	lru.data.MoveToFront(element)
	return entry.value, true
}

// Set 设置一个缓存值，如果缓存已满，则移除最久未使用的条目
func (lru *LRUCache) Set(key string, value interface{}, expiration time.Duration) {
	lru.lock.Lock()
	defer lru.lock.Unlock()

	if element, ok := lru.cache[key]; ok {
		lru.data.MoveToFront(element)
		entry := element.Value.(*CacheEntry)
		entry.value = value
		entry.expiration = expiration
		entry.lastUsed = time.Now()
		return
	}

	if lru.data.Len() >= lru.cap {
		lru.removeOldest()
	}

	entry := &CacheEntry{
		key:        key,
		value:      value,
		lastUsed:   time.Now(),
		expiration: expiration,
	}
	element := lru.data.PushFront(entry)
	lru.cache[key] = element
}

// removeElement 从缓存中移除一个条目
func (lru *LRUCache) removeElement(e *list.Element) {
	lru.data.Remove(e)
	kv := e.Value.(*CacheEntry)
	delete(lru.cache, kv.key)
}

// removeOldest 移除最久未使用的条目
func (lru *LRUCache) removeOldest() {
	e := lru.data.Back()
	if e != nil {
		lru.removeElement(e)
	}
}

func main() {
	// 示例
	lru := NewLRUCache(2)
	lru.Set("key1", "value1", 5*time.Second)
	lru.Set("key2", "value2", 2*time.Second)

	for e := lru.data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*CacheEntry).value)
	}
	fmt.Println("===========================")

	// 等待一会...
	time.Sleep(3 * time.Second)

	// "key2"应该已经过期
	if val, ok := lru.Get("key2"); ok {
		fmt.Println("key2:", val) // 这行不会被执行，因为key2已经过期
	} else {
		fmt.Println("key2 has expired")
	}

	// "key1"还在有效期内
	if val, ok := lru.Get("key1"); ok {
		fmt.Println("key1:", val) // 输出: key1: value1
	} else {
		fmt.Println("key1 has expired")
	}
}
