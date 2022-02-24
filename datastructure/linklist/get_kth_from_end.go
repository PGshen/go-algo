package linklist

/**
 * 获取链表的导数第K个
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	// p1节点先走k步
	p1 := head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	// p1, p2节点同时走，知道p1到达终点
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
