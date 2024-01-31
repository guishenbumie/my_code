package main

import "math"

/* 接雨水 */

func main() {

}

// https://mp.weixin.qq.com/s/8E2WHPdArs3KwSwaxFunHw
func trap(height []int) int {
	left, right := 0, len(height)-1
	lMax, rMax := 0, 0
	ans := 0
	for left < right {
		lMax = int(math.Max(float64(lMax), float64(height[left])))
		rMax = int(math.Max(float64(rMax), float64(height[right])))

		if lMax < rMax {
			ans += lMax - height[left]
			left++
		} else {
			ans += rMax - height[right]
			right--
		}
	}
	return ans
}
