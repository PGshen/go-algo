package main

/**
不同的二叉搜索树
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
 */
func numTrees(n int) int {
	// 备忘录
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	return numTreeRange(1, n, memo)
}

func numTreeRange(low, high int, memo [][]int) int {
	if low > high {
		// base case
		return 1
	}
	if memo[low][high] != 0 {
		return memo[low][high]
	}
	res := 0
	// 以mid为根节点，总数为左子树个数与右子树个数的乘积
	for mid := low; mid < high; mid++ {
		numLow := numTreeRange(low, mid-1, memo)
		numHigh := numTreeRange(mid+1, high, memo)
		res += numLow * numHigh
	}
	memo[low][high] = res
	return res
}


/**
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
 */
func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return []*TreeNode{}
	}
	return generate(1, n)
}

func generate(low, high int) []*TreeNode {
	var res []*TreeNode
	if low > high {
		res = append(res, nil)
	}
	for mid := low; mid < high; mid++ {
		leftTrees := generate(low, mid-1)
		rightTrees := generate(mid+1, high)
		// 左右子树叉乘
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				res = append(res, &TreeNode{
					Val:   mid,
					Left:  leftTree,
					Right: rightTree,
				})
			}
		}
	}
	return res
}

