package main

import "math"

/* 戳气球 */

func main() {

}

func maxCoins(nums []int) int {
	n := len(nums)
	points := []int{1}
	points = append(points, nums...)
	points = append(points, 1)
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	for i := n; i >= 0; i-- { //i从下往上
		for j := i + 1; j < n+2; j++ { //j从左到右
			for k := i + 1; k < j; k++ { //戳破的气球
				dp[i][j] = int(math.Max(
					float64(dp[i][j]),
					float64(dp[i][k]+dp[k][j]+points[i]*points[j]*points[k]),
				))
			}
		}
	}
	return dp[0][n+1]
}
