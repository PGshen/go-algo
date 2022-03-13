package linklist

import (
	"testing"
)

func Test_sortList(t *testing.T) {
	listNode := &ListNode{3, &ListNode{99, &ListNode{22, &ListNode{21, &ListNode{Val: 39}}}}}
	printListNode(listNode)
	printListNode(mergeSortList(listNode))
}
