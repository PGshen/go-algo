package graph

func isBipartite(graph [][]int) bool {
	length := len(graph)
	var (
		ok      bool   = true
		color   []bool = make([]bool, length)
		visited []bool = make([]bool, length)
	)
	var traverse func([][]int, int)
	traverse = func(graph [][]int, v int) {
		if !ok {
			return
		}
		visited[v] = true
		for _, w := range graph[v] {
			if !visited[w] {
				color[w] = !color[v]
				traverse(graph, w)
			} else {
				if color[w] == color[v] {
					ok = false
				}
			}
		}
	}
	for i := 0; i < length; i++ {
		if !visited[i] {
			traverse(graph, i)
		}
	}
	return ok
}
