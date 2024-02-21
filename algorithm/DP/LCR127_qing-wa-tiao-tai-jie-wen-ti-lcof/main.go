package main

/* 跳跃训练 */

// 和70题基本一样
func main() {

}

func trainWays(num int) int {
	if num == 0 {
		return 1
	}
	memo := make([]int, num+1)
	var dp func(n int) int
	dp = func(n int) int {
		if n <= 2 {
			return n
		}
		if memo[n] > 0 {
			return memo[n]
		}
		memo[n] = (dp(n-1) + dp(n-2)) % 1000000007
		return memo[n]
	}
	return dp(num)
}
