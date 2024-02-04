package main

import "fmt"

/* 不同的二叉搜索树 */

func main() {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(4))
	fmt.Println(numTrees(5))
}

//func numTrees(n int) int {
//	G := make([]int, n+1)
//	G[0], G[1] = 1, 1
//	for i := 2; i <= n; i++ {
//		for j := 1; j <= i; j++ {
//			G[i] += G[j-1] * G[i-j]
//		}
//	}
//	return G[n]
//}

func numTrees(n int) int {
	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n+1)
	}

	var count func(lo, hi int) int
	count = func(lo, hi int) int {
		if lo > hi {
			return 1
		}
		if memo[lo][hi] != 0 {
			return memo[lo][hi]
		}
		res := 0
		for mid := lo; mid <= hi; mid++ {
			left := count(lo, mid-1)
			right := count(mid+1, hi)
			res += left * right
		}
		memo[lo][hi] = res
		return res
	}

	return count(1, n)
}
