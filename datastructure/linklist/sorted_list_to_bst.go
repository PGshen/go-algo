package linklist

/**
 * 有序链表转为平衡二叉树
 * 思路： 链表转为数组，分治+递归
 */
func sortedListToBST(head *ListNode) *TreeNode {
	arr := []int{}
	for head != nil { // 转为数组
		arr = append(arr, head.Val)
		head = head.Next
	}
	return buildBST(arr, 0, len(arr)-1)
}

func buildBST(arr []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	// 分治递归
	mid := (left + right) >> 1
	root := &TreeNode{Val: arr[mid]}
	root.Left = buildBST(arr, left, mid-1)
	root.Right = buildBST(arr, mid+1, right)
	return root
}
