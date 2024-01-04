package main

import "fmt"

/* 岛屿数量 */

func main() {
	//visited := make([][]bool, 3)
	//for i := 0; i < 3; i++ {
	//	visited[i] = make([]bool, 2)
	//}
	//fmt.Println(visited)
	fmt.Println(numIslands([][]byte{
		{49, 49},
		{48, 48},
	}))
}

func dfs(grid [][]byte, i, j int) {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

func numIslands(grid [][]byte) int {
	var res int
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++           //发现陆地，数量+1
				dfs(grid, i, j) //以该陆地为中心，进行遍历，且将该陆地连接的陆地都变成水，不然需要维护一个visited对象，麻烦
			}
		}
	}
	return res
}
