package tree

/**
 * 判断是否合法的平衡二叉树
 */
func isValidBST(root *TreeNode) bool {
	return validBST(root, nil, nil)
}

// 左子树比root小，右子树比root大
func validBST(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && min.Val > root.Val {
		return false
	}
	if max != nil && max.Val < root.Val {
		return false
	}
	return validBST(root.Left, min, root) && validBST(root.Right, root, max)
}

/**
 * 二叉搜索树的搜索操作
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else if root.Val > val {
		return searchBST(root.Left, val)
	}
	return nil
}

/**
 * 插入新节点，找到叶子位置然后插入
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

/**
 * 删除节点。找到要删除的节点，然后根据类型判断
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 取右子树最小节点
		minNode := getMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

func getMin(root *TreeNode) *TreeNode {
	p := root
	for p.Left != nil {
		p = p.Left
	}
	return p
}
