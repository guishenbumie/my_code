package main

import "math"

/* 最小路径和 */

func main() {

}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = int(math.Min(
				float64(dp[i-1][j]),
				float64(dp[i][j-1]),
			)) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

//
//func minPathSum(grid [][]int) int {
//	m, n := len(grid), len(grid[0])
//	memo := make([][]int, m)
//	for i := 0; i < m; i++ {
//		memo[i] = make([]int, n)
//		for j := 0; j < n; j++ {
//			memo[i][j] = -1
//		}
//	}
//	var dp func(i, j int) int
//	dp = func(i, j int) int {
//		if i == 0 && j == 0 {
//			return grid[0][0]
//		}
//		if i < 0 || j < 0 {
//			return math.MaxInt
//		}
//		if memo[i][j] != -1 {
//			return memo[i][j]
//		}
//		memo[i][j] = int(math.Min(
//			float64(dp(i-1, j)),
//			float64(dp(i, j-1)),
//		)) + grid[i][j]
//		return memo[i][j]
//	}
//	return dp(m-1, n-1)
//}
