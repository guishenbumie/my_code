package main

import "fmt"

/* 子集 */

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		ans = append(ans, append([]int(nil), path...))
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
