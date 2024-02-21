package main

import (
	"fmt"
	"sort"
)

/* 最长递增子序列 */

func main() {
	arr := []int{1, 2, 3, 5}
	fmt.Println(sort.SearchInts(arr, 2))
	fmt.Println(sort.SearchInts(arr, 4))
	fmt.Println(sort.SearchInts(arr, 6))
	//fmt.Println(lengthOfLIS([]int{10, 9, 2}))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18, 1}))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}

//func lengthOfLIS(nums []int) int {
//	arr := make([]int, len(nums))
//	for i := 0; i < len(arr); i++ {
//		arr[i] = 1
//	}
//	for i := 0; i < len(nums); i++ {
//		for j := 0; j < i; j++ {
//			if nums[i] > nums[j] {
//				if arr[i] < arr[j]+1 {
//					arr[i] = arr[j] + 1
//				}
//			}
//		}
//	}
//	res := math.MinInt32
//	for i := 0; i < len(arr); i++ {
//		if arr[i] > res {
//			res = arr[i]
//		}
//	}
//	return res
//}

// 使用二分，这种时间复杂度更小
func lengthOfLIS(nums []int) int {
	arr := make([]int, 0)
	for _, x := range nums {
		j := sort.SearchInts(arr, x) //通过二分查找，查找到元素的位置
		if j == len(arr) {           //>=x的arr[j]不存在
			arr = append(arr, x)
		} else {
			arr[j] = x
		}
	}
	return len(arr)
}
