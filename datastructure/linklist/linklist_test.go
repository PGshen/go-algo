package linklist

import (
	"fmt"
	"testing"
)

// 测试合并两个有序链表
func TestMergeTwoLists(t *testing.T) {
	listNode1 := &ListNode{1, &ListNode{4, &ListNode{9, &ListNode{Val: 20}}}}
	listNode2 := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}

	listNode := mergeTwoLists1(listNode1, listNode2)
	printListNode(listNode)
}

// 测试合并k个有序链表
func TestMergeKLists(t *testing.T) {
	listNode1 := &ListNode{1, &ListNode{4, &ListNode{9, &ListNode{Val: 20}}}}
	listNode2 := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}
	listNode3 := &ListNode{2, &ListNode{5, &ListNode{Val: 99}}}

	l1 := mergeKLists([]*ListNode{listNode1, listNode2, listNode3})
	printListNode(l1)

	listNode1 = &ListNode{1, &ListNode{4, &ListNode{9, &ListNode{Val: 20}}}}
	listNode2 = &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}
	listNode3 = &ListNode{2, &ListNode{5, &ListNode{Val: 99}}}
	l2 := mergeKLists2([]*ListNode{listNode1, listNode2, listNode3})
	printListNode(l2)

}

// 测试 获取链表倒数第k个节点
func TestGetKthFromEnd(t *testing.T) {
	listNode1 := &ListNode{1, &ListNode{4, &ListNode{9, &ListNode{Val: 20}}}}
	printListNode(listNode1)
	printListNode(getKthFromEnd(listNode1, 2))
}

// 测试 删除链表倒数第k个节点
func TestRemoveKthFromEnd(t *testing.T) {
	listNode1 := &ListNode{1, &ListNode{4, &ListNode{9, &ListNode{Val: 20}}}}
	printListNode(listNode1)
	printListNode(removeKthFromEnd(listNode1, 2))
}

// 测试获取链表的中间节点
func TestMiddleNode(t *testing.T) {
	listNode2 := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}
	printListNode(listNode2)
	printListNode(middleNode(listNode2))
}

// 测试是否有环
func TestHasCycle(t *testing.T) {
	listNode2 := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}
	printListNode(listNode2)
	fmt.Println(hasCycle(listNode2))
}

// 测试检测环的起点
func TestDetectCycle(t *testing.T) {
	listNode1 := &ListNode{Val: 52}
	listNode2 := &ListNode{66, &ListNode{Val: 67, Next: listNode1}}
	listNode1.Next = listNode2
	listNode3 := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39, Next: listNode1}}}}}
	printNode(detectCycle(listNode3))
	listNod4 := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}
	printNode(detectCycle(listNod4))
}

// 测试获取两个链表的相交点
func TestGetInterSectionNode(t *testing.T) {
	listNode1 := &ListNode{1, &ListNode{Val: 2}}
	listNode2 := &ListNode{3, &ListNode{Val: 4, Next: listNode1}}
	listNode3 := &ListNode{5, &ListNode{6, &ListNode{Val: 7, Next: listNode1}}}
	printNode(getIntersectionNode(listNode2, listNode3))
}

// 测试反转链表
func TestReverseLinklist(t *testing.T) {
	listNode := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}
	printListNode(listNode)
	printListNode(reverseLinklist(listNode))
	printListNode(reverseLinklist2(listNode))
}

// 测试反转链表前N个节点
func TestReverseLinklistN(t *testing.T) {
	listNode := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}
	printListNode(listNode)
	printListNode(reverseLinklistN(listNode, 3))
	// printListNode(reverseLinklistN2(listNode, 3))
	listNode2 := &ListNode{Val: 1, Next: nil}

	printListNode(reverseLinklistN2(listNode2, 1))
}

// 测试反转链表中间部分呢的节点
func TestReverseLinklistBetween(t *testing.T) {
	listNode := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}
	printListNode(listNode)
	printListNode(reverseLinklistBetween(listNode, 2, 4))
	printListNode(reverseLinklistBetween2(listNode, 2, 4))
}

// 测试 k 个一组反转链表
func TestReverseKGroup(t *testing.T) {
	listNode := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}
	printListNode(listNode)
	printListNode(reverseKGroup(listNode, 2))
}

// 测试是否回文链表
func TestIsPalindrome(t *testing.T) {
	listNode := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{23, &ListNode{Val: 39}}}}}
	fmt.Println(isPalindrome(listNode))
	listNode2 := &ListNode{3, &ListNode{8, &ListNode{22, &ListNode{8, &ListNode{Val: 3}}}}}
	fmt.Println(isPalindrome(listNode2))
	listNode3 := &ListNode{3, &ListNode{8, &ListNode{8, &ListNode{Val: 3}}}}
	fmt.Println(isPalindrome(listNode3))

}

// 测试两个链表相加
func TestAddTwoNumbers(t *testing.T) {
	listNode1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{Val: 9}}}}}}}
	listNode2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{Val: 9}}}}
	printListNode(addTwoNumbers(listNode1, listNode2))
}
