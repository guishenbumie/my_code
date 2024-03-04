package main

import "fmt"

/* 连续数组 */

// 和525题一样
func main() {
	fmt.Println(1 ^ 1)
	fmt.Println(1 ^ 2)
}

func findMaxLength(nums []int) int {
	m := map[int]int{0: -1}
	ans := 0
	count := 0
	for i, x := range nums {
		if x == 1 {
			count++
		} else {
			count--
		}
		if preIdx, ok := m[count]; ok {
			if i-preIdx > ans {
				ans = i - preIdx
			}
		} else {
			m[count] = i
		}
	}
	return ans
}
