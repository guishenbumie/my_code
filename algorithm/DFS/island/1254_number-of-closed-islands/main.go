package main

/* 统计封闭岛屿的数目 */

func main() {

}

func closedIsland(grid [][]int) int {
	var res int
	m := len(grid)
	n := len(grid[0])
	//和200题不同的是，对于边界的岛屿不算在封闭岛屿内，因此先把边界的岛屿淹掉，再统计
	//将左右的边界岛屿淹掉
	for i := 0; i < m; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, n-1)
	}
	//将上下的边界岛屿淹掉
	for j := 0; j < n; j++ {
		dfs(grid, 0, j)
		dfs(grid, m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == 1 {
		//已经是海水了
		return
	}
	grid[i][j] = 1
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
