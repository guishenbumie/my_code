package main

import "sort"

/* 三数之和 */

// 和15题一样
func main() {

}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		tuples := twoSum(nums, 0-nums[i], i+1)
		for _, tuple := range tuples {
			ans = append(ans, append(tuple, nums[i]))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

func twoSum(nums []int, target, start int) [][]int {
	res := [][]int{}
	lo, hi := start, len(nums)-1
	for lo < hi {
		left := nums[lo]
		right := nums[hi]
		sum := left + right
		if sum < target {
			for lo < hi && nums[lo] == left {
				lo++
			}
		} else if sum > target {
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else {
			res = append(res, []int{left, right})
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		}
	}
	return res
}
