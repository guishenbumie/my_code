package main

import "fmt"

/* 判断子序列 */

func main() {
	fmt.Println(isSubsequence("b", "abc"))
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
