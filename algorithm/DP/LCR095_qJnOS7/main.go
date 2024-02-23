package main

import "math"

/* 最长公共子序列 */

// 和1143题一样
func main() {

}

func longestCommonSubsequence(text1 string, text2 string) int {
	n1 := len(text1)
	n2 := len(text2)
	c1 := []byte(text1)
	c2 := []byte(text2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if c1[i-1] == c2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(
					float64(dp[i-1][j]),
					float64(dp[i][j-1]),
				))
			}
		}
	}
	return dp[n1][n2]
}
