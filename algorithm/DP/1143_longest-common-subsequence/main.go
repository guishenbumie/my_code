package main

import (
	"fmt"
	"math"
)

/* 最长公共子序列 */

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
}

//func longestCommonSubsequence(text1 string, text2 string) int {
//	var dp func(s1 []byte, i int, s2 []byte, j int) int
//	dp = func(s1 []byte, i int, s2 []byte, j int) int {
//		if i == len(s1) || j == len(s2) {
//			return 0
//		}
//		if s1[i] == s2[j] {
//			return 1 + dp(s1, i+1, s2, j+1)
//		} else {
//			temp1 := dp(s1, i+1, s2, j)
//			temp2 := dp(s1, i, s2, j+1)
//			if temp1 > temp2 {
//				return temp1
//			} else {
//				return temp2
//			}
//		}
//	}
//	return dp([]byte(text1), 0, []byte(text2), 0)
//}

//func longestCommonSubsequence(text1 string, text2 string) int {
//	memo := make([][]int, len(text1))
//	for i := 0; i < len(memo); i++ {
//		memo[i] = make([]int, len(text2))
//		for j := 0; j < len(text2); j++ {
//			memo[i][j] = -1
//		}
//	}
//	var dp func(s1 []byte, i int, s2 []byte, j int) int
//	dp = func(s1 []byte, i int, s2 []byte, j int) int {
//		if i == len(s1) || j == len(s2) {
//			return 0
//		}
//		if memo[i][j] != -1 {
//			return memo[i][j]
//		}
//		if s1[i] == s2[j] {
//			memo[i][j] = 1 + dp(s1, i+1, s2, j+1)
//		} else {
//			temp1 := dp(s1, i+1, s2, j)
//			temp2 := dp(s1, i, s2, j+1)
//			if temp1 > temp2 {
//				memo[i][j] = temp1
//			} else {
//				memo[i][j] = temp2
//			}
//		}
//		return memo[i][j]
//	}
//	return dp([]byte(text1), 0, []byte(text2), 0)
//}

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	c1 := []byte(text1)
	c2 := []byte(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
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
	return dp[m][n]
}
