package main

/* 所有可能的路径 */

// 和797题一样
func main() {

}

func allPathsSourceTarget(graph [][]int) [][]int {
	ans := [][]int{}
	stk := []int{0}
	var dfs func(x int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			ans = append(ans, append([]int(nil), stk...))
			return
		}
		for _, y := range graph[x] {
			stk = append(stk, y)
			dfs(y)
			stk = stk[:len(stk)-1]
		}
	}
	dfs(0)
	return ans
}
