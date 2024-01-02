package main

import (
	"fmt"
	"strings"
)

/* 验证回文字符串 */

//如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
//
//字母和数字都属于字母数字字符。
//
//给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
//
//
//
//示例 1：
//
//输入: s = "A man, a plan, a canal: Panama"
//输出：true
//解释："amanaplanacanalpanama" 是回文串。
//示例 2：
//
//输入：s = "race a car"
//输出：false
//解释："raceacar" 不是回文串。
//示例 3：
//
//输入：s = " "
//输出：true
//解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
//由于空字符串正着反着读都一样，所以是回文串。

func main() {
	//fmt.Println([]byte("0"))
	//fmt.Println([]byte("9"))
	//fmt.Println([]byte("a"))
	//fmt.Println([]byte("z"))
	//fmt.Println([]byte("A"))
	//fmt.Println([]byte("Z"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	ss := []byte(s)
	arr := []byte("")
	for _, v := range ss {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
			arr = append(arr, v)
		}
	}
	//fmt.Println(string(arr))
	arrr := strings.ToLower(string(arr))
	for i := 0; i < len(arrr)/2; i++ {
		if arrr[i] != arrr[len(arrr)-i-1] {
			return false
		}
	}
	return true

	//ss := []byte(s)
	//arr := []byte("")
	//for _, v := range ss {
	//	if (v >= 48 && v <= 57) || (v >= 97 && v <= 122) {
	//		arr = append(arr, v)
	//	}
	//	if v >= 65 && v <= 90 {
	//		arr = append(arr, v+32)
	//	}
	//}
	//fmt.Println(string(arr))
	//for i := 0; i < len(arr)/2; i++ {
	//	if arr[i] != arr[len(arr)-i-1] {
	//		return false
	//	}
	//}
	//return true
}
