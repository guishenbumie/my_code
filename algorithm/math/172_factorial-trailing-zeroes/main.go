package main

import "fmt"

/* 阶乘后的零 */

func main() {
	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
	fmt.Println(trailingZeroes(0))
}

func trailingZeroes(n int) int {
	var res int
	for d := n; d/5 > 0; d = d / 5 {
		res += d / 5
	}
	return res
}
