package linklist

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

// 复制带随机指针的链表
// 迭代复制
func copyRandomList(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}
	// 复制链表
	p := head
	for p != nil {
		newNode := &RandomNode{Val: p.Val, Next: p.Next}
		p.Next = newNode
		p = newNode.Next
	}
	// 复制随机指针，画图就很清晰了
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}
	// 解除关联
	newRandomList := head.Next
	p = head
	for p != nil {
		tmp := p.Next
		p.Next = tmp.Next
		p = tmp
	}
	return newRandomList
}

// 复制带随机指针的链表
// 回溯 + hash
func copyRandomList2(head *RandomNode) *RandomNode {
	if head == nil {
		return head
	}
	hash := map[*RandomNode]*RandomNode{}
	return dfs(head, hash)
}

func dfs(node *RandomNode, hash map[*RandomNode]*RandomNode) *RandomNode {
	if node == nil {
		return nil
	}
	if v, ok := hash[node]; ok {
		return v
	}
	clone := &RandomNode{Val: node.Val}
	hash[node] = clone
	clone.Next = dfs(node.Next, hash)
	clone.Random = dfs(node.Random, hash)
	return clone
}
