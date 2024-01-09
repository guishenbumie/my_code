package main

import (
	"fmt"
	"sort"
)

/* 三数之和 */

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		tuples := twoSum(nums, 0-nums[i], i+1)
		for _, tuple := range tuples {
			res = append(res, append(tuple, nums[i]))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSum(nums []int, target, start int) [][]int {
	res := make([][]int, 0)
	lo, hi := start, len(nums)-1
	for lo < hi {
		left := nums[lo]
		right := nums[hi]
		sum := left + right
		if sum == target {
			res = append(res, []int{left, right})
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else if sum < target {
			for lo < hi && nums[lo] == left {
				lo++
			}
		} else if sum > target {
			for lo < hi && nums[hi] == right {
				hi--
			}
		}
	}
	return res
}
