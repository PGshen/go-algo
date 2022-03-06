package linklist

/**
 * 链表排序
 * 归并排序
 */
func mergeSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 分治
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmp := slow.Next
	slow.Next = nil
	left := mergeSortList(head)
	right := mergeSortList(tmp)
	// 合并
	return mergeTwoLists2(left, right)
}

// 插入排序
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: -5000}
	for head != nil {
		p1 := dummy
		for p1.Next != nil && p1.Next.Val < head.Val { // 找到比当前节点大一点的节点的前置节点
			p1 = p1.Next
		}
		node := &ListNode{Val: head.Val}
		node.Next = p1.Next
		p1.Next = node
		head = head.Next
	}
	return dummy.Next
}
