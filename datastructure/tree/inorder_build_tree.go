package main

import "fmt"

// 根据前序+中序遍历恢复二叉树
// 思路：前序遍历的第一个值为根节点，然后找到这个值在中序遍历的索引，这个索引的左侧为左子树，右侧为右子树的节点，然后递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	index := 0
	for i, _ := range inorder {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}
	return &TreeNode{Val: preorder[0],
		Left:  buildTree(preorder[1:index+1], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
}

// 根据后序+中序遍历恢复二叉树
// 思路：前序遍历的最后一个值为根节点，然后找到这个值在中序遍历的索引，这个索引的左侧为左子树，右侧为右子树的节点，然后递归
func buildTree2(postorder []int, inorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	index := 0
	for i, _ := range inorder {
		if inorder[i] == postorder[len(postorder)-1] {
			index = i
			break
		}
	}
	return &TreeNode{Val: postorder[len(postorder)-1],
		Left:  buildTree2(postorder[:index], inorder[:index]),
		Right: buildTree2(postorder[index:len(postorder)-1], inorder[index+1:]),
	}
}

func main() {
	//preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9,15,7,20,3}
	//node := buildTree(preorder, inorder)
	node := buildTree2(postorder, inorder)
	fmt.Printf("%v", node)
}
