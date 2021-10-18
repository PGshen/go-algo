package main

/**
所有可能的路径
给一个有n个结点的有向无环图，找到所有从0到n-1的路径并输出（不要求按顺序）

二维数组的第 i 个数组中的单元都表示有向图中 i 号结点所能到达的下一些结点

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-paths-from-source-to-target
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		res  = [][]int{}
		n    = len(graph)
		path = []int{}
	)
	var traverse func(int, []int)
	traverse = func(s int, path []int) {
		path = append(path, s)
		if s == n-1 {
			res = append(res, append([]int{}, path...))
			path = path[:len(path)-1]
			return
		}
		for _, node := range graph[s] {
			traverse(node, path)
		}
		path = path[:len(path)-1]
	}
	traverse(0, path)
	return res
}

