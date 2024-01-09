package main

import "fmt"

/* 错误的集合 */

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{1, 1}))
}

func findErrorNums(nums []int) []int {
	var sum int
	m := make(map[int]struct{})
	var idxv int
	var tempSum int
	for k, v := range nums {
		tempSum = tempSum + k + 1
		sum += v
		if _, ok := m[v]; ok {
			idxv = v
		} else {
			m[v] = struct{}{}
		}
	}
	idxvv := (tempSum - sum) + idxv
	return []int{idxv, idxvv}
}
