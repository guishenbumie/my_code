package main

import "fmt"

/* 在排序数组中查找元素的第一个和最后一个位置 */

//给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回 [-1, -1]。
//
//你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
//
//
//示例 1：
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//示例 2：
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//示例 3：
//
//输入：nums = [], target = 0
//输出：[-1,-1]
//
//
//提示：
//
//0 <= nums.length <= 105
//-109 <= nums[i] <= 109
//nums 是一个非递减数组
//-109 <= target <= 109

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}

func searchRange(nums []int, target int) []int {
	left := leftBound(nums, target)
	if left < 0 {
		return []int{-1, -1}
	}
	right := rightBound(nums, target)
	return []int{left, right}
}

func leftBound(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid - 1
		}
	}
	if left < 0 || left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

//func leftBound(nums []int, target int) int {
//	if len(nums) <= 0 {
//		return -1
//	}
//	left, right := 0, len(nums)
//	for left < right {
//		mid := left + (right-left)/2
//		if nums[mid] < target {
//			left = mid + 1
//		} else if nums[mid] > target {
//			right = mid
//		} else {
//			right = mid
//		}
//	}
//	return left
//}

func rightBound(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right < 0 || right >= len(nums) || nums[right] != target {
		return -1
	}
	return right
}

//func rightBound(nums []int, target int) int {
//	if len(nums) <= 0 {
//		return -1
//	}
//	left, right := 0, len(nums)
//	for left < right {
//		mid := left + (right-left)/2
//		if nums[mid] < target {
//			left = mid + 1
//		} else if nums[mid] > target {
//			right = mid
//		} else {
//			left = mid + 1
//		}
//	}
//	return left - 1
//}
