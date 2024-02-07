package main

/* 组合总和 */

// 和39题一样
func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	path := []int{}
	sum := 0
	var dfs func(i int)
	dfs = func(i int) {
		if target == sum {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		if sum > target {
			return
		}
		for j := i; j < len(candidates); j++ {
			sum += candidates[j]
			path = append(path, candidates[j])
			dfs(j)
			path = path[:len(path)-1]
			sum -= candidates[j]
		}
	}
	dfs(0)
	return ans
}
