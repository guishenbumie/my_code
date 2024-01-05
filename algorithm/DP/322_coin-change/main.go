package main

import (
	"fmt"
)

/* 零钱兑换 */

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
}

//func coinChange(coins []int, amount int) int {
//	if amount == 0 {
//		return 0
//	}
//	if amount < 0 {
//		return -1
//	}
//	res := math.MaxInt32
//	for _, c := range coins {
//		sub := coinChange(coins, amount-c)
//		if sub == -1 {
//			continue
//		}
//		if res > sub+1 {
//			res = sub + 1
//		}
//	}
//	if res == math.MaxInt32 {
//		return -1
//	} else {
//		return res
//	}
//}

//func coinChange(coins []int, amount int) int {
//	memo := make([]int, amount+1)
//	for i := 0; i < len(memo); i++ {
//		memo[i] = -666
//	}
//	var dp func(coins []int, amount int) int
//	dp = func(coins []int, amount int) int {
//		if amount == 0 {
//			return 0
//		}
//		if amount < 0 {
//			return -1
//		}
//		if memo[amount] != -666 { //备忘录中已经计算过了，直接返回
//			return memo[amount]
//		}
//		res := math.MaxInt32
//		for _, c := range coins {
//			sub := dp(coins, amount-c)
//			if sub == -1 {
//				continue
//			}
//			if res > sub+1 {
//				res = sub + 1
//			}
//		}
//		if res == math.MaxInt32 {
//			res = -1
//		}
//		memo[amount] = res
//		return res
//	}
//	return dp(coins, amount)
//}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, c := range coins {
			if i-c < 0 {
				continue
			}
			if dp[i] > 1+dp[i-c] {
				dp[i] = 1 + dp[i-c]
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
