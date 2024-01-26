package main

import "sort"

/* 小于k的两数之和 */

// 给你一个整数数组nums和整数k,返回最大和sum,满足存在i<j使得nums[i]+nums[j] = sum
// 且sum<k。如果没有满足此等式的i,j存在,则返回-1。
func main() {

}

func twoSumLessThanK(nums []int, k int) int {
	sort.Ints(nums)
	lo, hi := 0, len(nums)-1
	ans := -1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum >= k {
			hi--
		} else {
			if sum > ans {
				ans = sum
			}
			lo++
		}
	}
	return ans
}
