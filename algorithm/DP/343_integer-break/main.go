package main

import "math"

/* 整数拆分 */

func main() {

}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		max := 0
		for j := 1; j < i; j++ {
			max = int(math.Max(
				float64(max),
				math.Max(
					float64(j*(i-j)),
					float64(j*dp[i-j]),
				),
			))
		}
		dp[i] = max
	}
	return dp[n]
}
