package tree

import "math"

func countNodes(root *TreeNode) int {
	l, r := root, root
	lh, rh := 0, 0
	for l != nil {
		lh++
		l = l.Left
	}
	for r != nil {
		rh++
		r = r.Right
	}
	if lh == rh {
		return int(math.Pow(2, float64(lh))) - 1
	}
	// 左右子树总会有一个是完全二叉树
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
