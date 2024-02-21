package main

import "math"

/* 买卖股票的最佳时机含手续费 */

// 122升级版
func main() {

}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2) //j表示当前手底股票的状态，0表示没有股票，1表示持有股票
		if i == 0 {            //第一天，j=0肯定没收益，j=1收益就是扣去股票的钱
			dp[i][0] = 0
			dp[i][1] = -prices[i] - fee
			continue
		}
		//当前没有持有股票
		dp[i][0] = int(math.Max(
			float64(dp[i-1][0]),           //前一天没买卖，维持没有股票状态
			float64(dp[i-1][1]+prices[i]), //前一天把手底的股票卖了
		))
		//当前持有股票
		dp[i][1] = int(math.Max(
			float64(dp[i-1][1]),               //前一天没买卖，维持持有股票状态
			float64(dp[i-1][0]-prices[i]-fee), //前一天买入了股票
		))
	}
	return dp[n-1][0] //最后肯定是手底下没有股票的状态收益最大
}
