package main

import "fmt"

/* 编辑距离 */

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	c1 := []byte(word1)
	c2 := []byte(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if c1[i-1] == c2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else { //dp[i-1][j-1] 表示替换操作，dp[i-1][j] 表示删除操作，dp[i][j-1] 表示插入操作
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m][n]
}

func min(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		} else {
			return b
		}
	} else {
		if a > c {
			return c
		} else {
			return a
		}
	}
}
