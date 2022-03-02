package linklist

import "container/heap"

// 实现最小堆
type ListNodes []*ListNode

func (l *ListNodes) Len() int {
	return len(*l)
}

func (l *ListNodes) Less(i, j int) bool {
	return (*l)[i].Val < (*l)[j].Val
}

func (l *ListNodes) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *ListNodes) Pop() interface{} {
	n := len(*l)
	x := (*l)[n-1]
	*l = (*l)[:n-1]
	return x
}

func (l *ListNodes) Push(x interface{}) {
	*l = append(*l, x.(*ListNode))
}

// 合并k个有序链表
// 思路： 借助堆实现
func mergeKLists(lists []*ListNode) *ListNode {
	listNodes := &ListNodes{}
	heap.Init(listNodes)
	for _, v := range lists {
		if v != nil {
			heap.Push(listNodes, v)
		}
	}
	head := &ListNode{}
	idx := head
	for listNodes.Len() > 0 {
		val := heap.Pop(listNodes).(*ListNode)
		idx.Next = val
		if val.Next != nil {
			heap.Push(listNodes, val.Next)
		}
		idx = idx.Next
	}
	return head.Next
}

// 合并k个有序链表
// 思路，分治+归并
func mergeKLists2(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	mid := length >> 1                                                          //分治
	return mergeTwoLists1(mergeKLists2(lists[:mid]), mergeKLists2(lists[mid:])) //归并
}
