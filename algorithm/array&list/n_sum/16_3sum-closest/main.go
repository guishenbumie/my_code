package main

import (
	"math"
	"sort"
)

/* 最接近的三数之和 */

func main() {

}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)

	diff := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		sum := nums[i] + twoSumClosest(nums, target-nums[i], i+1)
		if math.Abs(float64(diff)) > math.Abs(float64(target-sum)) {
			diff = target - sum
		}
	}
	return target - diff
}

func twoSumClosest(nums []int, target, start int) int {
	lo, hi := start, len(nums)-1
	diff := math.MaxInt32
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if math.Abs(float64(diff)) > math.Abs(float64(target-sum)) {
			diff = target - sum
		}
		if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return target - diff
}
