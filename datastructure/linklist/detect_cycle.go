package linklist

/*
 * 检测链表是否有环
 */
func detectCycle(head *ListNode) *ListNode {
	// 快慢指针找到相遇点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	// 没有相遇点
	if fast == nil || fast.Next == nil {
		return nil
	}
	// slow回到头然后同步向前走
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
