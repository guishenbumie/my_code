package main

import "math"

/* 长度最小的子数组 */

// 和209题一样
func main() {

}

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	ans := math.MaxInt32
	sum := 0
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < ans {
				ans = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++

	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
