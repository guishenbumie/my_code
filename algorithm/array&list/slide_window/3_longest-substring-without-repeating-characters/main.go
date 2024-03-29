package main

import (
	"fmt"
	"math"
)

/* 无重复字符的最长子串 */

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//提示：
//
//0 <= s.length <= 5 * 104
//s 由英文字母、数字、符号和空格组成

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	window := map[byte]int{}
	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++

		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		//diff := right - left
		//if res < diff {
		//	res = diff
		//}
		res = int(math.Max(float64(res), float64(right-left)))
	}
	return res
}

//func lengthOfLongestSubstring(s string) int {
//	windows := map[byte]int{}
//	left, right := 0, 0
//	res := 0
//	for right < len(s) {
//		c := s[right]
//		right++
//		windows[c]++
//
//		for windows[c] > 1 {
//			d := s[left]
//			left++
//			windows[d]--
//		}
//
//		diff := right - left
//		if diff > res {
//			res = diff
//		}
//	}
//	return res
//}
