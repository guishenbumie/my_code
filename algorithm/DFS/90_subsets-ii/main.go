package main

import (
	"fmt"
	"sort"
)

/* 子集 II */

//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
//解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
//
//
//示例 1：
//
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
//示例 2：
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//提示：
//
//1 <= nums.length <= 10
//-10 <= nums[i] <= 10

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(nums) //先给数组排序，后面遍历的时候相同的数值仅遍历一次
	var dfs func(i int)
	dfs = func(i int) {
		ans = append(ans, append([]int(nil), path...))
		if i == len(nums) {
			return
		}
		for j := i; j < len(nums); j++ {
			if j > i && nums[j] == nums[j-1] {
				continue
			}
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
