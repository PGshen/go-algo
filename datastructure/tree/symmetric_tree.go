package tree

// 判断是否是对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs00(root.Left, root.Right)
}

func dfs00(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return dfs00(left.Left, right.Right) && dfs00(left.Right, right.Left)
}
