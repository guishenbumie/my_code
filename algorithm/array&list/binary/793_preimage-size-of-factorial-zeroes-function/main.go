package main

import "math"

/* 阶乘函数后 K 个零 */

func main() {

}

func preimageSizeFZF(k int) int {
	return rightBound(k) - leftBound(k) + 1
}

func leftBound(target int) int {
	left, right := 0, math.MaxInt
	for left < right {
		mid := left + (right-left)/2
		if trailingZeroes(mid) < target {
			left = mid + 1
		} else if trailingZeroes(mid) > target {
			right = mid
		} else {
			right = mid
		}
	}
	return left
}

func rightBound(target int) int {
	left, right := 0, math.MaxInt
	for left < right {
		mid := left + (right-left)/2
		if trailingZeroes(mid) < target {
			left = mid + 1
		} else if trailingZeroes(mid) > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left - 1
}

func trailingZeroes(n int) int {
	var res int
	for d := n; d/5 > 0; d = d / 5 {
		res += d / 5
	}
	return res
}
