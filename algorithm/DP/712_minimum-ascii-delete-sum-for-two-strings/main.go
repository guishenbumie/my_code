package main

import "fmt"

/* 两个字符串的最小ASCII删除和 */

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
	fmt.Println(minimumDeleteSum("delete", "leet"))
}

func minimumDeleteSum(s1 string, s2 string) int {
	memo := make([][]int, len(s1))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(s2))
		for j := 0; j < len(s2); j++ {
			memo[i][j] = -1
		}
	}
	var dp func(s1 []byte, i int, s2 []byte, j int) int
	dp = func(s1 []byte, i int, s2 []byte, j int) int {
		var res int
		if i == len(s1) {
			for ; j < len(s2); j++ {
				res += int(s2[j])
			}
			return res
		}
		if j == len(s2) {
			for ; i < len(s1); i++ {
				res += int(s1[i])
			}
			return res
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if s1[i] == s2[j] {
			memo[i][j] = dp(s1, i+1, s2, j+1)
		} else {
			temp1 := int(s1[i]) + dp(s1, i+1, s2, j)
			temp2 := int(s2[j]) + dp(s1, i, s2, j+1)
			if temp1 > temp2 {
				memo[i][j] = temp2
			} else {
				memo[i][j] = temp1
			}
		}
		return memo[i][j]
	}
	return dp([]byte(s1), 0, []byte(s2), 0)
}
