package main

import "math"

/* 砍竹子 I */

// 和343题一样
func main() {

}

func cuttingBamboo(bamboo_len int) int {
	dp := make([]int, bamboo_len+1)
	for i := 2; i <= bamboo_len; i++ {
		max := 0
		for j := 1; j < i; j++ {
			max = int(math.Max(
				float64(max),
				math.Max(
					float64(j*(i-j)),
					float64(j*dp[i-j]),
				),
			))
			dp[i] = max
		}
	}
	return dp[bamboo_len]
}
