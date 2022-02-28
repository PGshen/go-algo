package linklist

/**
 * 删除倒数第k个节点
 */
func removeKthFromEnd(head *ListNode, k int) *ListNode {
	dummy := ListNode{Next: head}
	x := getKthFromEnd(&dummy, k+1) // 倒数第n+1个节点。使用虚拟指针避免空情况
	x.Next = x.Next.Next
	return dummy.Next
}
