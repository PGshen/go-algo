package tree

/**
 * 思路，利用二叉搜索树的中序遍历符合升序来找对应的错误位置的元素
 * 第一个错误元素取前一个（父节点） 第二个错误元素取后一个（自身），然后交换值
 */
var parent, x, y *TreeNode

func recoverTree(root *TreeNode) {
	x, y, parent = nil, nil, nil
	inOrder(root)
	x.Val, y.Val = y.Val, x.Val
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)

	if parent != nil && parent.Val > root.Val {
		y = root
		if x == nil {
			x = parent
		} else {
			return
		}
	}
	parent = root

	inOrder(root.Right)
}
