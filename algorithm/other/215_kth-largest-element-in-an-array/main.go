package main

import (
	"fmt"
	"math/rand"
)

/* 数组中的第K个最大元素 */

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

// 获取第k个大的数
func findKthLargest(nums []int, k int) int {
	shuffle(nums) //防止数组已经是排序过了，出现快排的最糟糕的时间复杂度，先将数组随机打乱
	TopKSplit(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

// 获取前k大的数
func findKthLargestNums(nums []int, k int) []int {
	shuffle(nums) //防止数组已经是排序过了，出现快排的最糟糕的时间复杂度，先将数组随机打乱
	TopKSplit(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k:]
}

// 获取第k个小的数
func findKthSmallest(nums []int, k int) int {
	shuffle(nums) //防止数组已经是排序过了，出现快排的最糟糕的时间复杂度，先将数组随机打乱
	TopKSplit(nums, k, 0, len(nums)-1)
	return nums[k-1]
}

// 获取前k小的数
func findKthSmallestNums(nums []int, k int) []int {
	shuffle(nums) //防止数组已经是排序过了，出现快排的最糟糕的时间复杂度，先将数组随机打乱
	TopKSplit(nums, k, 0, len(nums)-1)
	return nums[:k]
}

func shuffle(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		nums[i], nums[r] = nums[r], nums[i]
	}
}

// 使用快排
func Parition(nums []int, start, stop int) int {
	if start > stop {
		return -1
	}
	pivot := nums[start]
	left, right := start, stop
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] < pivot {
			left++
		}
		nums[right] = nums[left]
	}
	//循环结束，left和right相等
	//确定基础元素pivot在数组中的位置
	nums[left] = pivot
	return left
}

func TopKSplit(nums []int, k, start, stop int) {
	if start < stop {
		idx := Parition(nums, start, stop)
		if idx == k {
			return
		} else if idx < k {
			TopKSplit(nums, k, idx+1, stop)
		} else if idx > k {
			TopKSplit(nums, k, start, idx-1)
		}
	}
}
