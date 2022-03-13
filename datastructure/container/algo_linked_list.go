package container

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	val  int
	Next *ListNode
	Prev *ListNode
}

type MyLinkedList struct {
	size int
	head *ListNode
	tail *ListNode
}

func NewLinklist() MyLinkedList {
	myLinkedList := MyLinkedList{
		size: 0,
		head: &ListNode{},
		tail: &ListNode{},
	}
	myLinkedList.head.Next = myLinkedList.tail
	myLinkedList.tail.Prev = myLinkedList.head
	return myLinkedList
}

func (this *MyLinkedList) Get(index int) int {
	node := this.getNode(index)
	if node == nil {
		return -1
	}
	return node.val
}

func (this *MyLinkedList) getNode(index int) *ListNode {
	if index < 0 || index >= this.size {
		return nil
	}
	p := this.head
	if index+1 < this.size-index {
		// 从头遍历
		for i := 0; i <= index; i++ {
			p = p.Next
		}
	} else {
		// 从尾遍历
		p = this.tail
		for i := this.size - 1; i >= index; i-- {
			p = p.Prev
		}
	}
	return p
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{val: val, Next: this.head.Next}
	this.head.Next.Prev = newNode
	newNode.Prev = this.head
	this.head.Next = newNode
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{val: val, Next: this.tail}
	newNode.Prev = this.tail.Prev
	this.tail.Prev.Next = newNode
	this.tail.Prev = newNode
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index >= this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	indexNode := this.getNode(index)
	newNode := &ListNode{val: val}
	newNode.Next = indexNode
	newNode.Prev = indexNode.Prev
	indexNode.Prev.Next = newNode
	indexNode.Prev = newNode
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	indexNode := this.getNode(index)
	indexNode.Next.Prev = indexNode.Prev
	indexNode.Prev.Next = indexNode.Next
	indexNode = nil
	this.size--
}

func (this *MyLinkedList) Show() {
	p := this.head.Next
	for p.Next != nil {
		fmt.Print(strconv.Itoa(p.val) + " ")
		p = p.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
