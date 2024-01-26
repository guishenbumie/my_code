package main

import "sort"

/* 较小的三数之和 */

// 给定一个长度为n的整数数组和一个目标值target,寻找能够使条件nums[i] + nums[j] + nums[k]
// <target成立的三元组i,j,k个数(0<=i<j<k<n)。
func main() {

}

func threeSumSmaller(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums)-2; i++ {
		count += twoSumSmaller(nums, target-nums[i], i+1)
	}
	return count
}

func twoSumSmaller(nums []int, target, start int) int {
	lo, hi := start, len(nums)-1
	count := 0
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum < target {
			count += hi - lo
			lo++
		} else {
			hi--
		}
	}
	return count
}
