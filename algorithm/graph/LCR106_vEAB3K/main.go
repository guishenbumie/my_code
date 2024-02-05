package main

/* 判断二分图 */

// 和785题一样
func main() {

}

func isBipartite(graph [][]int) bool {
	ans := true
	n := len(graph)
	color := make([]bool, n)
	visited := make([]bool, n)

	var traverse func(x int)
	traverse = func(x int) {
		if !ans {
			return
		}
		visited[x] = true
		for _, y := range graph[x] {
			if !visited[y] {
				color[y] = !color[x]
				traverse(y)
			} else {
				if color[y] == color[x] {
					ans = false
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			traverse(i)
		}
	}
	return ans
}
