package main

/* 组合 */

// 和77题一样
func main() {

}

func combine(n int, k int) [][]int {
	ans := [][]int{}
	path := []int{}
	var dfs func(i int)
	dfs = func(i int) {
		if len(path) == k {
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
