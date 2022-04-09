package tree

// 404. 左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	// 判断左叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}
	return res + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
