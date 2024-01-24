package main

import "math"

/* 绝对差不超过限制的最长连续子数组 */

func main() {

}

func longestSubarray(nums []int, limit int) int {
	ans := 0
	minQ, maxQ := []int{}, []int{}

	left, right := 0, 0
	for right < len(nums) {
		x := nums[right]

		//minQ队列，单调递增
		for len(minQ) > 0 && minQ[len(minQ)-1] > x {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, x)

		//maxQ队列，单调递减
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < x {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, x)

		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}
		ans = int(math.Max(float64(ans), float64(right-left+1)))

		right++
	}

	return ans
}
