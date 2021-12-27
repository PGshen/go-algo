package linklist

import "container/heap"

/**
单链表操作
*/

/**
1. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/

// 递归法
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		// l1的头节点后接 l1.Next l2排序后的链表
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists1(l1, l2.Next)
		return l2
	}
}

// 迭代法
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	// 移动curr指针
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	// 将剩余的接上
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return dummy.Next
}

/**
2. 合并k个有序链表
*/

// 基于堆
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

func mergeKLists(lists []*ListNode) *ListNode {
	listNodes := &ListNodes{}
	heap.Init(listNodes)
	// 初始化入堆
	for _, v := range lists {
		if v != nil {
			heap.Push(listNodes, v)
		}
	}
	head := &ListNode{}
	idx := head
	for listNodes.Len() > 0 {
		// 每次弹出堆顶
		val := heap.Pop(listNodes).(*ListNode)
		idx.Next = val
		if val.Next != nil {
			heap.Push(listNodes, val.Next)
		}
		idx = idx.Next
	}
	return head.Next
}
