package linklist

import (
	"testing"
)

// 测试合并两个有序链表
func TestMergeTwoLists(t *testing.T) {
	listNode1 := &ListNode{1, &ListNode{4, &ListNode{9, &ListNode{Val: 20}}}}
	listNode2 := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}

	listNode := mergeTwoLists1(listNode1, listNode2)
	printNode(listNode)
}

// 测试合并k个有序链表
func TestMergeKLists(t *testing.T) {
	listNode1 := &ListNode{1, &ListNode{4, &ListNode{9, &ListNode{Val: 20}}}}
	listNode2 := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}
	listNode3 := &ListNode{2, &ListNode{5, &ListNode{Val: 99}}}

	listNode := mergeKLists([]*ListNode{listNode1, listNode2, listNode3})
	printNode(listNode)
}

// 测试 获取链表倒数第k个节点
func TestGetKthFromEnd(t *testing.T) {
	listNode1 := &ListNode{1, &ListNode{4, &ListNode{9, &ListNode{Val: 20}}}}
	printNode(listNode1)
	printNode(getKthFromEnd(listNode1, 2))
}
