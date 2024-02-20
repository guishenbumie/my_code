package main

/* 爬楼梯 */

func main() {

}

// func climbStairs(n int) int {
//     if n <= 1 {
// 		return 1
// 	}
//     res := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		if i <= 1 {
// 			res[i] = 1
// 			continue
// 		}
// 		res[i] = res[i-1] + res[i-2]
// 	}
// 	return res[n-1] + res[n-2]
// }

func climbStairs(n int) int {
	memo := make([]int, n+1)
	var dp func(n int) int
	dp = func(n int) int {
		if n <= 2 {
			return n
		}
		if memo[n] > 0 {
			return memo[n]
		}
		memo[n] = dp(n-1) + dp(n-2)
		return memo[n]
	}
	return dp(n)
}
