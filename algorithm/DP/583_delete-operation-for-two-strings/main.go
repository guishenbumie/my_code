package main

import (
	"fmt"
	"math"
)

/* 两个字符串的删除操作 */

func main() {
	fmt.Println(minDistance("sea", "eat"))
	fmt.Println(minDistance("leetcode", "etco"))
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	c1 := []byte(word1)
	c2 := []byte(word2)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if c1[i-1] == c2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Max(
					float64(dp[i-1][j]),
					float64(dp[i][j-1]),
				))
			}
		}
	}
	return m - dp[m][n] + n - dp[m][n]
}
