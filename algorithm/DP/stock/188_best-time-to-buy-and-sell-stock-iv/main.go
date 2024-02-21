package main

import "math"

/* 买卖股票的最佳时机 IV */

// 123升级版
func main() {

}

func maxProfit(k int, prices []int) int {
	max_k, n := k, len(prices)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, max_k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		for k := max_k; k >= 1; k-- { //加入了次数的限制
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = int(math.Max(
				float64(dp[i-1][k][0]),
				float64(dp[i-1][k][1]+prices[i]),
			))
			dp[i][k][1] = int(math.Max(
				float64(dp[i-1][k][1]),
				float64(dp[i-1][k-1][0]-prices[i]),
			))
		}
	}
	return dp[n-1][max_k][0]
}
