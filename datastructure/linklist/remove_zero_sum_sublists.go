package linklist

/**
 * 将给定链表中和为0的子链表全部删除，最后得到的链表节点值之和应当不为0.
 * 思路，计算前缀和
 * https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/solution/qian-zhui-he-dan-lian-biao-de-qian-zhui-kwr05/
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	sumMap := map[int]*ListNode{}
	dummy := &ListNode{0, head}
	sum := 0
	sumMap[0] = dummy
	p := head
	for p != nil {
		sum += p.Val
		if node, ok := sumMap[sum]; ok {
			delNode := node.Next
			node.Next = p.Next
			// 删除中间的映射
			delSum := sum
			for delNode != p {
				delSum += delNode.Val
				delete(sumMap, delSum)
				delNode = delNode.Next
			}
		} else {
			sumMap[sum] = p
		}
		p = p.Next
	}
	return dummy.Next
}
