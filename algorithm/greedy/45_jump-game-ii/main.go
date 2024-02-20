package main

import "math"

/* 跳跃游戏 II */

func main() {

}

// 贪心算法
func jump(nums []int) int {
	end, farthest, jumps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = int(math.Max(float64(nums[i]+i), float64(farthest)))
		if end == i {
			jumps++
			end = farthest
		}
	}
	return jumps
}
