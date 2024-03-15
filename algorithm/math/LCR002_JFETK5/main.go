package main

import (
	"math"
	"strconv"
)

/* 二进制求和 */

// 和67题一样
func main() {

}

func addBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)
	n := int(math.Max(
		float64(lenA),
		float64(lenB),
	))
	ans := ""
	carry := 0
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
