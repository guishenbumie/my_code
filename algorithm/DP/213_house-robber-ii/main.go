package main

import "math"

/* 打家劫舍 II */

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

//
//func rob(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//
//	//三种情况
//	//1：收尾都不抢，这种情况肯定最小，可以不看
//	//2：首抢，尾不抢
//	//3：首不抢，尾抢
//	memo1 := make([]int, len(nums))
//	memo2 := make([]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		memo1[i] = -1
//		memo2[i] = -1
//	}
//
//	return int(math.Max(
//		float64(dp(nums, memo1, 0, len(nums)-2)),
//		float64(dp(nums, memo2, 1, len(nums)-1)),
//	))
//}
//
//func dp(nums, memo []int, start, end int) int {
//	if start > end {
//		return 0
//	}
//	if memo[start] != -1 {
//		return memo[start]
//	}
//	memo[start] = int(math.Max(
//		float64(dp(nums, memo, start+1, end)),
//		float64(nums[start]+dp(nums, memo, start+2, end)),
//	))
//	return memo[start]
//}
