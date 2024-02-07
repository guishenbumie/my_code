package main

import "strconv"

/* 目标和 */

// 和494题一样
func main() {

}

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	memo := map[string]int{}
	var dp func(i, remain int) int
	dp = func(i, remain int) int {
		if i == len(nums) {
			if remain == 0 {
				return 1
			}
			return 0
		}
		key := strconv.Itoa(i) + "," + strconv.Itoa(remain)
		if val, ok := memo[key]; ok {
			return val
		}
		result := dp(i+1, remain-nums[i]) + dp(i+1, remain+nums[i])
		memo[key] = result
		return result
	}
	return dp(0, target)
}
