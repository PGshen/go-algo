package linklist

import "container/list"

/**
 * 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
 * 初始状态下，所有 next 指针都被设置为 NULL。
 */

// 迭代法
func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	pre := root
	for pre.Left != nil {
		tmp := pre
		for tmp != nil { // 同一层，从左往右连接
			tmp.Left.Next = tmp.Right
			if tmp.Next != nil {
				tmp.Right.Next = tmp.Next.Left
			}
			tmp = tmp.Next
		}
		pre = pre.Left // 向下一层
	}
	return root
}

// 递归法
func connect2(root *Node) *Node {
	if root == nil || root.Left == nil {
		// 没有子节点
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil { // 父节点已连接，子节点也需要连接
		root.Right.Next = root.Next.Left
	}
	connect2(root.Left)
	connect2(root.Right)
	return root
}

// 层次遍历
func connect3(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		pre := queue.Remove(queue.Front()).(*Node) //出队列
		if pre.Left != nil {
			queue.PushBack(pre.Left)
		}
		if pre.Right != nil {
			queue.PushBack(pre.Right)
		}
		for i := 1; i < length; i++ {
			cur := queue.Remove(queue.Front()).(*Node)
			pre.Next = cur
			pre = cur
			if pre.Left != nil {
				queue.PushBack(pre.Left)
			}
			if pre.Right != nil {
				queue.PushBack(pre.Right)
			}
		}
	}
	return root
}

// 普通二叉树
func connect4(root *Node) *Node {
	if root == nil {
		return root
	}
	pre := root
	for pre.Left != nil || pre.Right != nil || pre.Next != nil {
		flag := false
		tmp := pre
		dummy := &Node{}
		for tmp != nil { // 同一层，从左往右连接
			if tmp.Left != nil {
				dummy.Next = tmp.Left
				dummy = dummy.Next
				if flag == false { // 确定最左子树
					pre = tmp.Left
					flag = true
				}
			}
			if tmp.Right != nil {
				dummy.Next = tmp.Right
				dummy = dummy.Next
				if flag == false {
					pre = tmp.Right
					flag = true
				}
			}
			tmp = tmp.Next
		}
		if flag == false { // 没有了
			break
		}
	}
	return root
}
