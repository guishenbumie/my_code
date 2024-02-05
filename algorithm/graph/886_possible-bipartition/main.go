package main

/* 可能的二分法 */

func main() {

}

func possibleBipartition(n int, dislikes [][]int) bool {
	color := make([]bool, n+1)
	visited := make([]bool, n+1)
	ok := true

	// 转化成邻接表表示图结构
	graph := buildGraph(n, dislikes)

	var traverse func(int)
	traverse = func(v int) {
		if !ok {
			return
		}
		visited[v] = true
		for _, w := range graph[v] {
			if !visited[w] {
				color[w] = !color[v]
				traverse(w)
			} else {
				if color[w] == color[v] {
					ok = false
				}
			}
		}
	}

	// 全图遍历，防止非连通图出现
	for v := 1; v <= n; v++ {
		if !visited[v] {
			traverse(v)
		}
	}

	return ok
}

// 建图函数
func buildGraph(n int, dislikes [][]int) [][]int {
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, 0)
	}
	for _, edge := range dislikes {
		v := edge[1]
		w := edge[0]
		// 「无向图」相当于「双向图」
		// v -> w
		graph[v] = append(graph[v], w)
		// w -> v
		graph[w] = append(graph[w], v)
	}
	return graph
}
