package linklist

/*
 * 反转链表，迭代法
 */
func reverseLinklist(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var cur = head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

/*
 * 反转链表，递归法
 * 定义函数返回反转后链表的最后一个节点
 */
func reverseLinklist2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseLinklist2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

/*
 * 反转链表的前N个节点,迭代法
 */
func reverseLinklistN(head *ListNode, n int) *ListNode {
	var pre *ListNode = nil
	var cur = head
	for cur != nil && n > 0 {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
		n--
	}
	head.Next = cur
	return pre
}

var successor *ListNode

/*
 * 反转链表的前N个节点，递归法
 */
func reverseLinklistN2(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseLinklistN2(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}

/*
 * 反转链表的M到N区间的节点，迭代法
 * 索引从1开始
 */
func reverseLinklistBetween(head *ListNode, m, n int) *ListNode {
	var pre *ListNode = nil
	var cur = head
	for cur != nil && m > 1 { // 前期移动m个节点
		pre = cur
		cur = cur.Next
		m--
		n--
	}
	p1 := pre
	p2 := cur
	for cur != nil && n > 0 { // n理解为反转的次数
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
		n--
	}
	p1.Next = pre
	p2.Next = cur
	return head
}

/*
 * 反转链表的M到N区间的节点，递归法
 * 索引从1开始
 */
func reverseLinklistBetween2(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return reverseLinklistN2(head, n)
	}
	head.Next = reverseLinklistBetween2(head.Next, m-1, n-1) // 向前推进
	return head
}

/*
 * k 个一组反转链表
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil { // 不足k个
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 反转链表a, b 之间的节点
func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode = nil
	var cur *ListNode = a
	for cur != b {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
