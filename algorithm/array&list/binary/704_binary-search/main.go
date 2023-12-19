package main

import "fmt"

/* 二分查找 */

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//
//
//示例 1:
//
//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//示例 2:
//
//输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1
//
//
//提示：
//
//你可以假设 nums 中的所有元素是不重复的。
//n 将在 [1, 10000]之间。
//nums 的每个元素都将在 [-9999, 9999]之间。

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(leftBound([]int{1, 2, 2, 2, 2}, 2))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
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

// 寻找左侧边界的二分查找
// 比如[1, 2, 2, 2, 2]，target=2，上面的就不适用
func leftBound(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 //搜索区间变为[mid+1,right]
		} else if nums[mid] > target {
			right = mid - 1 //搜索区间变为[left,mid-1]
		} else { //nums[mid]==target
			right = mid - 1 //收缩右侧边界，变为[left,mid-1]
		}
	}
	//检查出界情况，target比数组的max元素还大会出现这种情况
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// 寻找右侧边界的二分查找
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
	//检查出界情况，target比数组的min元素还小会出现这种情况
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
