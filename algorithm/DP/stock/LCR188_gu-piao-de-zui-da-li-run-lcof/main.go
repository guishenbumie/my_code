package main

import "math"

/* 买卖芯片的最佳时机 */

// 和121题一样
func main() {

}

func bestTiming(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = int(math.Max(
			float64(dp[i-1][0]),
			float64(dp[i-1][1]+prices[i]),
		))
		dp[i][1] = int(math.Max(
			float64(dp[i-1][1]),
			float64(-prices[i]),
		))
	}
	return dp[n-1][0]
}
