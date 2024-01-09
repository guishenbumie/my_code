package main

import (
	"fmt"
	"sort"
)

/* 两数之和 */

func main() {
	fmt.Println(twoSum2([]int{1, 3, 1, 2, 2, 3}, 4))
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		m[v] = k
	}
	for k, v := range nums {
		temp := target - v
		if k2, ok := m[temp]; ok && k2 != k {
			return []int{k, k2}
		}
	}
	return nil
}

// 题目升级版本
// 题目告诉我们可以假设nums中有且只有一个答案,且需要我我们返回对应元素的索引,现在修改这些条件:
// nums中可能有多对儿元素之和都等于target,请你的算法返回所有和为target的元素对儿,其中不能出现重复。
// 比如：比如说输入为nums=[1,3,1,2,2,3], target=4,那么算法返回的结果就是:[[1,3] , [2,2]]
func twoSum2(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	lo, hi := 0, len(nums)-1
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
