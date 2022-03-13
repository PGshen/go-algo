package linklist

import (
	"fmt"
	"math/rand"
	"strconv"
)

// 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTreeNode(treeNode *TreeNode) {
	if treeNode == nil {
		return
	}
	printTreeNode(treeNode.Left)
	fmt.Print(strconv.Itoa(treeNode.Val) + " -> ")
	printTreeNode(treeNode.Right)
	rand.Intn(1)
}
