package tree

/**
 * 二叉树的最近公共祖先
 * 这个函数的功能有三个：给定两个节点 pp 和 qq
 * 如果 pp 和 qq 都存在，则返回它们的公共祖先；
 * 如果只存在一个，则返回存在的一个；
 * 如果 pp 和 qq 都不存在，则返回NULL
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}
