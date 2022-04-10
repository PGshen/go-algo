package tree

import "strconv"

// 257. 二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	ret := []string{}
	if root == nil {
		return ret
	}
	var dfs func(*TreeNode, string)
	dfs = func(root *TreeNode, path string) {
		if path != "" {
			path = path + "->"
		}
		path = path + strconv.Itoa(root.Val)

		if root.Left == nil && root.Right == nil {
			ret = append(ret, path)
			return
		}
		if root.Left != nil {
			dfs(root.Left, path)
		}
		if root.Right != nil {
			dfs(root.Right, path)
		}
	}
	dfs(root, "")
	return ret
}
