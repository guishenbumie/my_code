package main

import (
	"fmt"
	"math"
)

/* 最长递增子序列 */

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}

func lengthOfLIS(nums []int) int {
	arr := make([]int, len(nums))
	for i := 0; i < len(arr); i++ {
		arr[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if arr[i] < arr[j]+1 {
					arr[i] = arr[j] + 1
				}
			}
		}
	}
	res := math.MinInt32
	for i := 0; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}
