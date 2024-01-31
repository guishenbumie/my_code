package main

import "strings"

/* 验证回文串 */

// 和125题一样
func main() {

}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	arr := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
			arr = append(arr, c)
		}
	}
	for i := 0; 2*i+1 < len(arr); i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}
