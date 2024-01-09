package main

/* 阶乘后的零 */

func main() {

}

func trailingZeroes(n int) int {
	var res int
	for d := n; d/5 > 0; d = d / 5 {
		res += d / 5
	}
	return res
}
