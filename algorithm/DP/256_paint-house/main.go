package main

import "math"

/* 粉刷房子 */

func main() {

}

func minCost(costs [][]int) int {
	m := len(costs)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 3)
	}
	//粉刷第一个房子的费用是固定的
	for j := 0; j < 3; j++ {
		dp[0][j] = costs[0][j]
	}
	//穷举
	for i := 1; i < m; i++ {
		dp[i][0] = int(math.Min(
			float64(dp[i-1][1]),
			float64(dp[i-1][2]),
		)) + costs[i][0]
		dp[i][1] = int(math.Min(
			float64(dp[i-1][0]),
			float64(dp[i-1][2]),
		)) + costs[i][1]
		dp[i][2] = int(math.Min(
			float64(dp[i-1][0]),
			float64(dp[i-1][1]),
		)) + costs[i][2]
	}
	ans := math.MaxInt32
	for j := 0; j < 3; j++ {
		ans = int(math.Min(
			float64(ans),
			float64(dp[m-1][j]),
		))
	}
	return ans
}
