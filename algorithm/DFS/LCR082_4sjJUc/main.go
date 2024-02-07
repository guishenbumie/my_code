package main

import "sort"

/* 组合总和 II */

// 和40题一样
func main() {

}

func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	sort.Ints(candidates)
	path := []int{}
	sum := 0
	var dfs func(i int)
	dfs = func(i int) {
		if sum == target {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		if sum > target {
			return
		}
		for j := i; j < len(candidates); j++ {
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}
			sum += candidates[j]
			path = append(path, candidates[j])
			dfs(j + 1)
			path = path[:len(path)-1]
			sum -= candidates[j]
		}
	}
	dfs(0)
	return ans
}
