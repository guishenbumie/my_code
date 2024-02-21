package main

import "math"

/* 地下城游戏 */

func main() {

}

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == m-1 && j == n-1 {
			if dungeon[i][j] >= 0 {
				return 1
			} else {
				return -dungeon[i][j] + 1
			}
		}
		if i == m || j == n {
			return math.MaxInt
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		res := -dungeon[i][j] + int(math.Min(
			float64(dp(i, j+1)),
			float64(dp(i+1, j)),
		))
		if res <= 0 {
			memo[i][j] = 1
		} else {
			memo[i][j] = res
		}
		return memo[i][j]
	}

	return dp(0, 0)
}
