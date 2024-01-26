package main

import "sort"

func main() {

}

func twoSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 2, target, 0)
}

func threeSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 3, target, 0)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 4, target, 0)
}

func nSum(nums []int, n, target, start int) [][]int {
	sz := len(nums)
	res := [][]int{}
	if n < 2 || sz < n { //至少是2sum，并且数组大小不应该小于n
		return res
	}
	if n == 2 {
		lo, hi := start, sz-1
		for lo < hi {
			left, right := nums[lo], nums[hi]
			sum := left + right
			if sum < target {
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else { //sum == target
				res = append(res, []int{left, right})
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else { //n>2的情况，递归计算(n-1)Sum的结果
		for i := start; i < sz; i++ {
			sub := nSum(nums, n-1, target-nums[i], i+1)
			for _, arr := range sub {
				arr = append(arr, nums[i])
				res = append(res, arr)
			}
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}
