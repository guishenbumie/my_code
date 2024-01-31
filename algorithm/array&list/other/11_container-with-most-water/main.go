package main

import "math"

/* 盛最多水的容器 */

func main() {

}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		h := int(math.Min(float64(height[left]), float64(height[right])))
		area := h * (right - left)
		if area > ans {
			ans = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}
