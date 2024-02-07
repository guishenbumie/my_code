package main

import (
	"strconv"
)

/* 不同岛屿的数量 */

func main() {

}

func numDistinctIslands(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	set := make(map[string]struct{})
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				str := dfs(grid, i, j, "", 666)
				set[str] = struct{}{}
			}
		}
	}
	return len(set)
}

func dfs(grid [][]int, i, j int, str string, dir int) string {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return ""
	}
	if grid[i][j] == 0 {
		return ""
	}
	//前序遍历位置，进入(i,j)
	grid[i][j] = 0
	str = str + strconv.Itoa(dir) + ","
	str = dfs(grid, i-1, j, str, 1) + dfs(grid, i+1, j, str, 2) + dfs(grid, i, j-1, str, 3) + dfs(grid, i, j+1, str, 4)
	//后序遍历位置，离开(i,j)
	str = str + strconv.Itoa(-dir) + ","
	return str
}
