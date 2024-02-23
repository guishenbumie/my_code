package main

import "math"

/* 零钱兑换 */

// 和322题一样
func main() {

}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = int(math.Min(
					float64(dp[i]),
					float64(dp[i-coins[j]]+1),
				))
			}
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
