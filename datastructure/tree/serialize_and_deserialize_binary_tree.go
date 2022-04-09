package tree

import (
	"strconv"
	"strings"
)

// 297. 二叉树的序列化与反序列化

type Codec struct {
}

func Constructor2() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	list := strings.Split(data, ",")
	return buildTree3(&list)
}

func buildTree3(list *[]string) *TreeNode {
	rootVal := (*list)[0]
	*list = (*list)[1:]
	if rootVal == "#" {
		return nil
	}
	Val, _ := strconv.Atoi(rootVal)
	root := &TreeNode{Val: Val}
	root.Left = buildTree3(list)
	root.Right = buildTree3(list)
	return root
}
