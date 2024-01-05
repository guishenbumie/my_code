package main

import (
	"fmt"
	"math"
)

/* 最大子数组和 */

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		temp := nums[i] + dp[i-1]
		if temp > nums[i] {
			dp[i] = temp
		} else {
			dp[i] = nums[i]
		}
	}

	res := math.MinInt32
	for i := 0; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
