package linklist

func isPalindrome(head *ListNode) bool {
	mid := getMid(head)
	left := head
	right := reverse(mid, nil)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

// 获取中间节点，如果是偶数的话返回后一个；如果是奇数则返回中点的后一位
func getMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		return slow.Next
	} else {
		return slow
	}
}
