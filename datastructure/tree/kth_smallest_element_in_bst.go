package tree

var res int  // 记录结果
var rank int // 记录次序

func kthSmallest(root *TreeNode, k int) int {
	rank = 0
	res = 0
	dfs(root, k)
	return res
}

// 中序遍历
func dfs(root *TreeNode, k int) {
	if root == nil {
		return
	}
	dfs(root.Left, k)
	rank++
	if rank == k {
		res = root.Val
		return
	}
	dfs(root.Right, k)
}
