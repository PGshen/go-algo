package main

import (
	"container/heap"
	"fmt"
)

// 实现heap的方法
type Heap []int

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func main_() {
	h := &Heap{}
	heap.Init(h)

	h.Push(5)
	h.Push(2)
	h.Push(1)
	h.Push(8)
	h.Push(4)
	h.Push(6)
	h.Push(2)
	fmt.Println(h)

	heap.Init(h)
	fmt.Println(h)

	heap.Remove(h, 0)
	fmt.Println(h)

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
}
