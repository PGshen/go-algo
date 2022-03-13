package linklist

type MultiNode struct {
	Val   int
	Prev  *MultiNode
	Next  *MultiNode
	Child *MultiNode
}

// 扁平化多级双向链表, 深度递归
// leetcode 430
func flattern(head *MultiNode) *MultiNode {
	dfs1(head)
	return head
}

func dfs1(head *MultiNode) *MultiNode {
	last := head
	for head != nil {
		if head.Child == nil {
			last = head
			head = head.Next
		} else {
			tmp := head.Next
			childLast := dfs1(head.Child)

			head.Next = head.Child
			head.Child.Prev = head
			head.Child = nil

			if tmp != nil {
				tmp.Prev = childLast
				childLast.Next = tmp
			}
			head = tmp
			last = childLast
		}
	}
	return last
}

func flattern2(head *MultiNode) *MultiNode {
	dummy := &MultiNode{Next: head}
	for head != nil {
		if head.Child != nil {
			tmp := head.Next
			child := head.Child

			head.Next = child
			child.Prev = head
			head.Child = nil

			last := head
			for last != nil {
				last = last.Next
			}
			if tmp != nil {
				tmp.Prev = last
				last.Next = tmp
			}
		}
		head = head.Next
	}
	return dummy.Next
}
