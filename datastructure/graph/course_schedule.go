package graph

/**
 * 207. 课程表
 * 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
 * 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
 * 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
 * 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
 */
func canFinish(numCourses int, prerequisites [][]int) bool {
	nodeCon := make(map[int][]int)
	inDegree := make([]int, numCourses)
	n := len(prerequisites)
	// 计算每个节点的入度,计算临接表
	for i := 0; i < n; i++ {
		fromNode := prerequisites[i][1] // 出节点
		toNode := prerequisites[i][0]   // 入节点
		inDegree[toNode]++              // 入度++
		if _, exist := nodeCon[fromNode]; exist {
			nodeCon[fromNode] = append(nodeCon[fromNode], toNode)
		} else {
			nodeCon[fromNode] = []int{toNode}
		}
	}
	// 入度为0的入队列
	var queue []int
	for node, indCnt := range inDegree {
		if indCnt == 0 {
			queue = append(queue, node)
		}
	}

	cnt := 0
	for len(queue) != 0 {
		selectedNode := queue[0] // 选中一门入度为0的课，出列
		queue = queue[1:]
		cnt++

		toNodes := nodeCon[selectedNode] // 获取被它依赖的节点所有节点
		if len(toNodes) != 0 {           // 有节点依赖于它
			for i := 0; i < len(toNodes); i++ {
				toNode := toNodes[i]
				inDegree[toNode]--         // 被依赖的节点入度-1
				if inDegree[toNode] == 0 { // 此节点依赖的所有节点都已经被选完
					queue = append(queue, toNode)
				}
			}
		}
	}
	return cnt == numCourses
}
