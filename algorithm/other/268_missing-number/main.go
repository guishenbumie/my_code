package main

import "sort"

/* 丢失的数字 */

func main() {

}

func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		if num != i {
			return i
		}
	}
	return len(nums)
}
