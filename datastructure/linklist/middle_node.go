package linklist

/**
 * 获取链表的中间节点
 * 思路： 快慢指针
 */
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
