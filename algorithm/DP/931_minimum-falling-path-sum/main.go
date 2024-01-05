package main

import (
	"fmt"
	"math"
)

/* 下降路径最小和 */

func main() {
	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
	fmt.Println(minFallingPathSum([][]int{{-19, 57}, {-40, -5}}))
}

func min(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		} else {
			return b
		}
	} else {
		if a > c {
			return c
		} else {
			return a
		}
	}
}

//func minFallingPathSum(matrix [][]int) int {
//	n := len(matrix)
//	res := math.MaxInt32
//	var dp func(matrix [][]int, i, j int) int
//	dp = func(matrix [][]int, i, j int) int {
//		if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
//			return 99999
//		}
//		if i == 0 {
//			return matrix[i][j]
//		}
//		return matrix[i][j] + min(dp(matrix, i-1, j-1), dp(matrix, i-1, j), dp(matrix, i-1, j+1))
//	}
//	for i := 0; i < n; i++ {
//		temp := dp(matrix, n-1, i)
//		if temp < res {
//			res = temp
//		}
//	}
//	return res
//}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	res := math.MaxInt32
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = 66666
		}
	}
	var dp func(matrix [][]int, i, j int) int
	dp = func(matrix [][]int, i, j int) int {
		if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
			return 99999
		}
		if i == 0 {
			return matrix[i][j]
		}
		if memo[i][j] != 66666 {
			return memo[i][j]
		}
		memo[i][j] = matrix[i][j] + min(dp(matrix, i-1, j-1), dp(matrix, i-1, j), dp(matrix, i-1, j+1))
		return memo[i][j]
	}
	for i := 0; i < n; i++ {
		temp := dp(matrix, n-1, i)
		if temp < res {
			res = temp
		}
	}
	return res
}
