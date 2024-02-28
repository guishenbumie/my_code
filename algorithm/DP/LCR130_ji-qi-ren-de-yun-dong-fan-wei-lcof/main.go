package main

import "fmt"

func main() {
	fmt.Println(wardrobeFinishing(4, 7, 5))
}

func wardrobeFinishing(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	res := 0
	dfs(m, n, k, 0, 0, visited, &res)
	return res
}

func dfs(m, n, k, i, j int, visited [][]bool, res *int) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return // 超出索引边界
	}
	if i/10+i%10+j/10+j%10 > k {
		return // 坐标和超出 k 的限制
	}
	if visited[i][j] {
		return // 之前已经访问过当前坐标
	}
	*res++
	visited[i][j] = true
	dfs(m, n, k, i+1, j, visited, res)
	dfs(m, n, k, i, j+1, visited, res)
	dfs(m, n, k, i-1, j, visited, res)
	dfs(m, n, k, i, j-1, visited, res)
}
