package tree

/**
 * 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和
 */
var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	traverse(root)
	return root
}

// 从右往左的中序遍历正好符合降序遍历
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Right)
	sum += root.Val
	root.Val = sum
	traverse(root.Left)
}
