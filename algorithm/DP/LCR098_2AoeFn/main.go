package main

/* 不同路径 */

// 和62题一样
func main() {

}

func uniquePaths(m int, n int) int {
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == 0 && j == 0 {
			return 1
		}
		if i < 0 || j < 0 {
			return 0
		}
		if memo[i][j] > 0 {
			return memo[i][j]
		}
		memo[i][j] = dp(i-1, j) + dp(i, j-1)
		return memo[i][j]
	}
	return dp(m-1, n-1)
}
