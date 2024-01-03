package main

import "fmt"

/* 组合 */

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
//你可以按 任何顺序 返回答案。
//
//
//
//示例 1：
//
//输入：n = 4, k = 2
//输出：
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]
//示例 2：
//
//输入：n = 1, k = 1
//输出：[[1]]
//
//
//提示：
//
//1 <= n <= 20
//1 <= k <= n

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if k == len(path) {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := i; j <= n; j++ {
			path = append(path, j)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ans
}
