package main

import (
	"sort"
)

/* 俄罗斯套娃信封问题 */

func main() {

}

func maxEnvelopes(envelopes [][]int) int {
	//对i(就是宽度)进行升序排列，再对j(高度)进行降序排列，最后对高度的一维数组进行lts求最值
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] > envelopes[j][0] {
			return false
		} else if envelopes[i][0] < envelopes[j][0] {
			return true
		} else {
			return envelopes[i][1] > envelopes[j][1]
		}
	})
	arr := make([]int, len(envelopes))
	for i := 0; i < len(arr); i++ {
		arr[i] = envelopes[i][1]
	}
	return lengthOfLIS(arr)
}

// 求数组中最长递增子序列的长度
func lengthOfLIS(nums []int) int {
	//res := 0
	//arr := make([]int, len(nums))
	//for i := 0; i < len(nums); i++ {
	//	curr := nums[i]
	//	left, right := 0, res
	//	for left < right {
	//		mid := left + (right-left)/2
	//		if arr[mid] >= curr {
	//			right = mid
	//		} else {
	//			left = mid + 1
	//		}
	//	}
	//	if left == res {
	//		res++
	//	}
	//	arr[left] = curr
	//}
	//return res
	arr := make([]int, 0)
	for _, x := range nums {
		j := sort.SearchInts(arr, x)
		if j == len(arr) {
			arr = append(arr, x)
		} else {
			arr[j] = x
		}
	}
	return len(arr)
}

// 用下面的方式会超时，需要用二分
//func lengthOfLIS(nums []int) int {
//	arr := make([]int, len(nums))
//	for i := 0; i < len(arr); i++ { //初始化，每个子数组的最长递增子序列最少是它本身，1
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
