package main

import "math"

func main() {

}

// 0-1背包问题
// 给你一个可装载重量为W的背包和N个物品, 每个物品有重量和价值两个属性。
// 其中第i个物品的重量为wt[i], 价值为val[i], 现在让你用这个背包装物品, 最多能装的价值是多少?
// 比如：N= 3，W = 4，wt = [2, 1, 3]，val = [4, 2, 3]
func knapsack(W, N int, wt, val []int) int {
	dp := make([][]int, N+1) //dp[i][w] 的定义如下：对于前 i 个物品，当前背包的容量为 w，这种情况下可以装的最大价值是 dp[i][w]
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-wt[i-1] < 0 {
				//不能进背包
				dp[i][w] = dp[i-1][w]
			} else {
				//装或者不装，择优
				dp[i][w] = int(math.Max(
					float64(dp[i-1][w]),
					float64(dp[i-1][w-wt[i-1]]+val[i-1]),
				))
			}
		}
	}
	return dp[N][W]
}
