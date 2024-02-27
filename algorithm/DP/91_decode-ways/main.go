package main

/* 解码方法 */

func main() {

}

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}
	for i := 2; i <= n; i++ {
		c := s[i-1]
		d := s[i-2]
		if '1' <= c && c <= '9' { //s[i]本身可以作为一个字母
			dp[i] += dp[i-1]
		}
		if d == '1' || d == '2' && c <= '6' { //s[i]和s[i-1]结合起来表示一个字母
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
