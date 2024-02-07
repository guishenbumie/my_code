package main

/* 子集 */

// 和78题一样
func main() {

}

func subsets(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
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
