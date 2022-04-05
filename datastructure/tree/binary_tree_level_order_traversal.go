package tree

import "container/list"

// 二叉树的层次遍历，借助队列完成
func levelOrder(root *TreeNode) [][]int {
	resList := [][]int{}
	if root == nil {
		return resList
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		len := queue.Len()
		tempList := []int{}
		for i := 0; i < len; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			tempList = append(tempList, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		resList = append(resList, tempList)
	}
	return resList
}
