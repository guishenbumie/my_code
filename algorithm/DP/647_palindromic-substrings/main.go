package main

import "math"

/* 回文子串 */

func main() {

}

func countSubstrings(s string) int {
	ans := 0
	n := len(s)
	t := "$#"
	for i := 0; i < n; i++ {
		t += string(s[i]) + "#"
	}
	n = len(t)
	t += "!"

	dp := make([]int, n)
	iMax, rMax, ans := 0, 0, 0
	for i := 1; i < n; i++ {
		// 初始化dp
		if i <= rMax {
			dp[i] = int(math.Min(
				float64(rMax-i+1),
				float64(dp[2*iMax-i]),
			))
		} else {
			dp[i] = 1
		}
		//中心扩展
		for t[i+dp[i]] == t[i-dp[i]] {
			dp[i]++
		}
		//动态维护max
		if i+dp[i]-1 > rMax {
			iMax = i
			rMax = i + dp[i] - 1
		}
		//统计答案
		ans += dp[i] / 2
	}
	return ans
}
