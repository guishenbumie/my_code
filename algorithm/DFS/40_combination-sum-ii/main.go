package main

import (
	"fmt"
	"sort"
)

/* 组合总和 II */

//给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的每个数字在每个组合中只能使用 一次 。
//
//注意：解集不能包含重复的组合。
//
//
//
//示例 1:
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//示例 2:
//
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//]
//
//
//提示:
//
//1 <= candidates.length <= 100
//1 <= candidates[i] <= 50
//1 <= target <= 30

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(candidates)
	var sum int
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
			sum -= candidates[j]
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
