package main

import "fmt"

/* 搜索旋转排序数组 */

func main() {
	fmt.Println(search([]int{1}, 1))
	fmt.Println(search([]int{2, 3, 1}, 2))
}

func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	var f func(left, right int) int
	f = func(left, right int) int {
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] < target {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				return mid
			}
		}
		return -1
	}

	idx := findMin(nums)
	var res int
	if target > nums[len(nums)-1] {
		res = f(0, idx-1)
	} else {
		res = f(idx, len(nums)-1)
	}
	return res
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
