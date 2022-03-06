package main

import (
	"container/list"
	"fmt"
)

// LRUCache 简单LRU缓存实现
type LRUCache struct {
	Capacity int
	Cache    map[int]*list.Element
	List     *list.List
}

// Node 节点
type Node struct {
	Key   int
	Value int
}

// Constructor 构造器
func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Cache:    map[int]*list.Element{},
		List:     list.New(),
	}
}

// Get 获取值，如无返回-1
func (lru *LRUCache) Get(key int) int {
	v, ok := lru.Cache[key]
	if !ok {
		fmt.Printf("Get %d: %d\n", key, -1)
		return -1
	}
	lru.List.MoveToFront(v)
	fmt.Printf("Get %d: %d\n", key, v.Value)
	return v.Value.(Node).Value
}

// Put 缓存键值对
func (lru *LRUCache) Put(key int, value int) {
	len := lru.List.Len()
	node := Node{Key: key, Value: value}
	if v, ok := lru.Cache[key]; ok {
		// 删除然后重新插入头部
		lru.List.Remove(v)
		element := lru.List.PushFront(node)
		lru.Cache[key] = element
	} else {
		if len == lru.Capacity {
			// 达到最大容量了，先删掉队尾元素
			element := lru.List.Back()
			lru.List.Remove(element)
			delete(lru.Cache, element.Value.(Node).Key)
		}
		element := lru.List.PushFront(node)
		lru.Cache[key] = element
	}
	fmt.Printf("Put %d: %d\n", key, value)
}
