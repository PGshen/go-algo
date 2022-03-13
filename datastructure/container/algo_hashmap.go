package container

import "container/list"

const base = 769

type MyHashMap struct {
	data []list.List
}

type entry struct {
	key   int
	value int
}

func NewMyHashMap() MyHashMap {
	return MyHashMap{
		data: make([]list.List, base),
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % base
}

func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	for p := this.data[h].Front(); p != nil; p = p.Next() {
		if et := p.Value.(entry); et.key == key {
			p.Value = entry{key, value}
			return
		}
	}
	this.data[h].PushBack(entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for p := this.data[h].Front(); p != nil; p = p.Next() {
		if et := p.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for p := this.data[h].Front(); p != nil; p = p.Next() {
		if et := p.Value.(entry); et.key == key {
			this.data[h].Remove(p)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
