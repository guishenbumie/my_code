package main

import "sort"

/* 全排列 II  */

// 和47题一样
func main() {

}

func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	path := []int{}
	visited := make([]bool, len(nums))
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := 0; j < len(nums); j++ {
			if visited[j] {
				continue
			}
			if j > 0 && nums[j] == nums[j-1] && !visited[j-1] {
				continue
			}
			path = append(path, nums[j])
			visited[j] = true
			dfs(i + 1)
			visited[j] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
