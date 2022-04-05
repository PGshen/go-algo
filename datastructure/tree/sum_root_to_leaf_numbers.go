package tree

/**
 * 129. 求根节点到叶节点数字之和
 */
func sumNumbers(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}
	var dfs func(root *TreeNode, val int)
	dfs = func(root *TreeNode, val int) {
		newVal := val*10 + root.Val
		if root.Left == nil && root.Right == nil {
			sum += newVal
			return
		}
		if root.Left != nil {
			dfs(root.Left, newVal)
		}
		if root.Right != nil {
			dfs(root.Right, newVal)
		}
	}
	dfs(root, 0)
	return sum
}
