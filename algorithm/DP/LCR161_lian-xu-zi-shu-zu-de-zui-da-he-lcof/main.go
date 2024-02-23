package main

import "math"

/* 连续天数的最高销售额 */

// 和53题一样
func main() {

}

func maxSales(sales []int) int {
	n := len(sales)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = sales[0]
	for i := 1; i < n; i++ {
		dp[i] = int(math.Max(
			float64(sales[i]),
			float64(dp[i-1]+sales[i]),
		))
	}
	ans := math.MinInt
	for i := 0; i < n; i++ {
		ans = int(math.Max(
			float64(ans),
			float64(dp[i]),
		))
	}
	return ans
}
