package HashTableLink

import (
	"errors"
	"math"
)

type HashTableLink struct {
	Table map[int]*List
	Size int
	Cap int
}

func (ht *HashTableLink) Get(key string) (interface{}, error) {
	index := ht.Pos(key)
	item, err := ht.Find(index, key)
	if item == nil {
		return " ", errors.New("not found")
	}
	return item.value, err
}

func (ht *HashTableLink) Put(key, value string) {
	index := ht.Pos(key)
	if ht.Table[index] == nil {
		ht.Table[index] = NewList() // 新建一个节点
	}
	item := &Element{key, value}
	data, err := ht.Find(index, key)
	if err != nil {
		ht.Table[index].Append(item)
		ht.Size++
	} else {
		data.value = value
	}
}

func (ht *HashTableLink) Del(key string) error {
	index := ht.Pos(key)
	list := ht.Table[index]
	var val *Element
	list.Each(func(node Node) {
		if node.Value.(*Element).key == key {
			val = node.Value.(*Element)
		}
	})
	if val == nil {
		return nil
	}
	ht.Size--
	return list.Remove(val)
}

func (ht *HashTableLink) Foreach(f func(item *Element)) {
	for k := range ht.Table {
		if ht.Table[k] != nil {
			ht.Table[k].Each(func(node Node) {
				f(node.Value.(*Element))
			})
		}
	}
}

func (ht *HashTableLink) Pos(s string) int {
	return ht.HashCode(s) % ht.Cap // 根据哈希值计算-- 除余法
}

func (ht *HashTableLink) Find(i int, key string) (*Element, error) {
	list := ht.Table[i]
	var val *Element
	list.Each(func(node Node) {
		if node.Value.(*Element).key == key {
			val = node.Value.(*Element)
		}
	})
	if val == nil {
		return nil, errors.New("not found")
	}
	return val, nil
}

func (ht *HashTableLink) HashCode(str string) int {
	hash := int32(0)
	for i := 0; i < len(str); i++ {
		hash = hash<<5 - hash + int32(str[i])
		hash &= hash // 哈希计算
	}
	return int(math.Abs(float64(hash)))
}

type Element struct {
	key string
	value interface{}
}

func NewHashTableLink(cap int) *HashTableLink {
	table := make(map[int]*List, cap)
	return &HashTableLink{table, 0, cap}
}

type HashTableLinkImple interface {
	Get(key string) (interface{},error)
	Put(key,value string)
	Del(key string)error
	Foreach(f func(item *Element))
	Pos(s string) int
	Find (i int,key string)(*Element,error)
	HashCode(str string)int
}


