package main

import (
	"fmt"
	"sort"
)

/* 四数之和 */

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		triples := threeSum(nums, target-nums[i], i+1)
		for _, triple := range triples {
			res = append(res, append(triple, nums[i]))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func threeSum(nums []int, target, start int) [][]int {
	res := make([][]int, 0)
	for i := start; i < len(nums); i++ {
		tuples := twoSum(nums, target-nums[i], i+1)
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
		} else if sum > target {
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else if sum < target {
			for lo < hi && nums[lo] == left {
				lo++
			}
		}
	}
	return res
}
