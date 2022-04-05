package tree

// 173. 二叉搜索树迭代器
// 借助栈实现迭代的复杂度的尽量均摊

// BSTIterator 迭代器
type BSTIterator struct {
	stack []*TreeNode
}

// Constructor 构造器
func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{}
	for node := root; node != nil; node = node.Left {
		iterator.stack = append(iterator.stack, node)
	}
	return iterator
}

// Next 获取下个元素
func (it *BSTIterator) Next() int {
	// 一路向左，入栈
	top := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]                  // 出栈
	for node := top.Right; node != nil; node = node.Left { // 存在右子树
		it.stack = append(it.stack, node)
	}
	return top.Val
}

// HasNext 是否还有下个
func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}
