package main

import "math"

/* 打家劫舍 II */

// 和213题一样
func main() {

}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}

	var f func(arr []int) int
	f = func(arr []int) int {
		dp := make([]int, len(arr))
		dp[0] = arr[0]
		dp[1] = int(math.Max(float64(arr[0]), float64(arr[1])))
		for i := 2; i < len(arr); i++ {
			dp[i] = int(math.Max(
				float64(dp[i-1]),
				float64(dp[i-2]+arr[i]),
			))
		}
		return dp[len(arr)-1]
	}

	ans := int(math.Max(
		float64(f(nums[:n-1])),
		float64(f(nums[1:])),
	))
	return ans
}
