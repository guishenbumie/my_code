package main

import "math"

/* 跳跃游戏 */

func main() {

}

func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = int(math.Max(float64(farthest), float64(i+nums[i])))
		if farthest <= i {
			return false
		}
	}
	return farthest >= len(nums)-1
}
