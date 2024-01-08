package main

import "fmt"

/* 零钱兑换 II */

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(change(3, []int{2}))
	fmt.Println(change(10, []int{10}))
	fmt.Println(change(0, []int{10}))
	fmt.Println(change(10, []int{}))
}

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1) //i=金币面值，j=总额
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[n][amount]
}
