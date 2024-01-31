package main

import (
	"strings"
)

/* 验证回文串 */

func main() {

}

func isPalindrome(s string) bool {
	ss := []byte(s)
	arr := []byte("")
	for _, v := range ss {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
			arr = append(arr, v)
		}
	}
	arrr := strings.ToLower(string(arr))
	for i := 0; i < len(arrr)/2; i++ {
		if arrr[i] != arrr[len(arrr)-i-1] {
			return false
		}
	}
	return true
}
