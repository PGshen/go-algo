package tree

import "container/list"

// 二叉树的锯齿层次遍历，借助队列完成
func zigzagLevelOrder(root *TreeNode) [][]int {
	resList := [][]int{}
	if root == nil {
		return resList
	}
	queue := list.New()
	queue.PushBack(root)
	level := 0
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
		if level%2 == 1 { // 奇数行
			// 反转一下
			for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
				tempList[i], tempList[j] = tempList[j], tempList[i]
			}
		}
		resList = append(resList, tempList)
		level++
	}
	return resList
}
