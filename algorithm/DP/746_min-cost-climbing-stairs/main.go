package main

import "math"

/* 使用最小花费爬楼梯 */

func main() {

}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = int(math.Min(
			float64(dp[i-1]+cost[i-1]),
			float64(dp[i-2]+cost[i-2]),
		))
	}
	return dp[n]
}
