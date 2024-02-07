package main

/* 全排列 */

// 和46题一样
func main() {

}

func permute(nums []int) [][]int {
	ans := [][]int{}
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
			visited[j] = true
			path = append(path, nums[j])
			dfs(i + 1)
			path = path[:len(path)-1]
			visited[j] = false
		}
	}
	dfs(0)
	return ans
}
