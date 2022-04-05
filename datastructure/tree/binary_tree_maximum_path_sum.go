package tree

/**
 * 二叉树中的最大路径和
 */
func maxPathSum(root *TreeNode) int {
	maxSum := -1001
	var dfs func(*TreeNode) int
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		left := dfs(tn.Left)
		right := dfs(tn.Right)
		// 包含当前节点的最大路径和
		tmp := tn.Val
		if left > 0 {
			tmp += left
		}
		if right > 0 {
			tmp += right
		}
		if tmp > maxSum {
			maxSum = tmp
		}
		// 当前为根节点能向上提供的最大路径和
		max := left
		if right > max {
			max = right
		}
		if max > 0 {
			max = max + tn.Val
		} else {
			max = tn.Val
		}
		if max > 0 {
			return max
		}
		return 0
	}
	dfs(root)
	return maxSum
}
