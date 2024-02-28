package main

import "math"

/* 三角形最小路径和 */

// 和120题一样
func main() {

}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		// 因为求最小值，所以全都初始化为极大值
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = int(math.Min(
				float64(dp[i-1][j]),
				float64(dp[i-1][j-1]),
			)) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for j := 0; j < len(dp[n-1]); j++ {
		if dp[n-1][j] < ans {
			ans = dp[n-1][j]
		}
	}
	return ans
}
