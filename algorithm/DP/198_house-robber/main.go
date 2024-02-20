package main

import "math"

/* 打家劫舍 */

func main() {

}

func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	//状态：索引；选择：不抢，去+1再判断抢不抢 或者 抢，跳过+1，去+2判断抢或者不抢
	var dp func(start int) int
	dp = func(start int) int {
		if start >= len(nums) {
			return 0
		}
		if memo[start] != -1 {
			return memo[start]
		}
		memo[start] = int(math.Max(
			float64(dp(start+1)),
			float64(nums[start]+dp(start+2))),
		)
		return memo[start]
	}

	return dp(0)
}
