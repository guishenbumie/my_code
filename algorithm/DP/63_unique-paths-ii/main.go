package main

/* 不同路径 II */

func main() {

}

// 自底向上，不太好理解，还是看下面的方式吧
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	// 数组索引偏移一位，dp[0][..] dp[..][0] 代表 obstacleGrid 之外
	// 定义：到达 obstacleGrid[i][j] 的路径条数为 dp[i-1][j-1]
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// base case：如果没有障碍物，起点到起点的路径条数就是 1
	dp[1][1] = 0
	if obstacleGrid[0][0] == 0 {
		dp[1][1] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				// 跳过 base case
				continue
			}
			if obstacleGrid[i-1][j-1] == 1 {
				// 跳过障碍物的格子
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	// 返回到达 obstacleGrid[m-1][n-1] 的路径数量
	return dp[m][n]
}

// //自顶向下
// func uniquePathsWithObstacles(obstacleGrid [][]int) int {
//     m, n := len(obstacleGrid), len(obstacleGrid[0])
//     memo := make([][]int, m)
//     for i:=0; i<m; i++ {
//         memo[i] = make([]int, n)
//     }
//     return dp(obstacleGrid, memo, m-1, n-1)
// }

// func dp(obstacleGrid, memo [][]int, i, j int) int {
//     m, n := len(obstacleGrid), len(obstacleGrid[0])
//     if i < 0 || i >= m || j < 0 || j >= n || obstacleGrid[i][j] == 1  {
//         return 0
//     }
//     if i == 0 && j == 0 {
//         return 1
//     }
//     if memo[i][j] > 0 {
//         return memo[i][j]
//     }
//     memo[i][j] = dp(obstacleGrid, memo, i-1, j) + dp(obstacleGrid, memo, i, j-1)
//     return memo[i][j]
// }
