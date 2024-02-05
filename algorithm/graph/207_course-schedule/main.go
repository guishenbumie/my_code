package main

/* 课程表 */

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	onPath := make([]bool, numCourses)
	visited := make([]bool, numCourses)
	hasCycle := false

	var traverse func(graph [][]int, x int)
	traverse = func(graph [][]int, x int) {
		if onPath[x] { //出现环
			hasCycle = true
		}
		if visited[x] || hasCycle {
			return
		}
		visited[x] = true
		onPath[x] = true
		for _, y := range graph[x] {
			traverse(graph, y)
		}
		onPath[x] = false //后序遍历，回退上一层，继续扫描
	}

	graph := build(numCourses, prerequisites)
	for i := 0; i < numCourses; i++ {
		traverse(graph, i)
	}

	return !hasCycle
}

func build(n int, prerequisites [][]int) [][]int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range prerequisites {
		from := edge[1]
		to := edge[0]
		graph[from] = append(graph[from], to)
	}
	return graph
}
