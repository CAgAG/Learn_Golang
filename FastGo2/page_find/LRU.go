package main

import (
	"container/list"
	"fmt"
)

// ======================================================
type LRU struct {
	cap   int
	cache map[int]*list.Element
	deque *list.List
}

type Item struct {
	key   int
	value int
}

func NewLRU(cap int) LRU {
	return LRU{cap: cap, cache: make(map[int]*list.Element), deque: list.New()}
}

func (lru *LRU) Get(key int) int {
	if ele, ok := lru.cache[key]; ok {
		lru.deque.MoveToFront(ele)
		return ele.Value.(*Item).value
	}
	return -1
}

func (lru *LRU) Put(key int, value int) {
	if ele, ok := lru.cache[key]; ok {
		ele.Value.(*Item).value = value
		lru.deque.MoveToFront(ele)
		return
	}

	if lru.cap == lru.deque.Len() {
		back := lru.deque.Back()
		if back != nil {
			lru.deque.Remove(back)
			delete(lru.cache, back.Value.(*Item).value)
		}
	}

	cur_ele := lru.deque.PushFront(&Item{key: key, value: value})
	lru.cache[key] = cur_ele
}

// ======================================================

func main() {
	lru := NewLRU(3)

	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)

	for e := lru.deque.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*Item).value)
	}
	fmt.Println("===========================")

	lru.Put(4, 4)

	for e := lru.deque.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*Item).value)
	}
	fmt.Println("===========================")

	lru.Put(2, 2)

	for e := lru.deque.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*Item).value)
	}
	fmt.Println("===========================")

}
