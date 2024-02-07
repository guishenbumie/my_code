package main

import "fmt"

/* 子集 */

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
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
//nums 中的所有元素 互不相同

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		ans = append(ans, append([]int(nil), path...)) //这里注意，新生成一个slice添加给ans，如果直接添加path，因为是引用，后面会跟着变
		if i == len(nums) {
			return
		}
		for j := i; j < len(nums); j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
