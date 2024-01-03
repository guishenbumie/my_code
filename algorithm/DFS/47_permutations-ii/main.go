package main

import (
	"fmt"
	"sort"
)

/* 全排列 II */

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
//示例 1：
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
//[1,2,1],
//[2,1,1]]
//示例 2：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//提示：
//
//1 <= nums.length <= 8
//-10 <= nums[i] <= 10

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(nums)
	used := make([]bool, len(nums))
	var dfs func(i int)
	dfs = func(i int) {
		if len(path) == len(nums) {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := 0; j < len(nums); j++ {
			if used[j] {
				continue
			}
			if j > 0 && nums[j] == nums[j-1] && !used[j-1] {
				continue
			}
			path = append(path, nums[j])
			used[j] = true
			dfs(j + 1)
			path = path[:len(path)-1]
			used[j] = false
		}
	}
	dfs(0)
	return ans
}
