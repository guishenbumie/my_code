package main

import (
	"fmt"
	"math"
)

/* 最小覆盖子串 */

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//
//
//注意：
//
//对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//示例 1：
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//示例 2：
//
//输入：s = "a", t = "a"
//输出："a"
//解释：整个字符串 s 是最小覆盖子串。
//示例 3:
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//提示：
//
//m == s.length
//n == t.length
//1 <= m, n <= 105
//s 和 t 由英文字母组成

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
}

func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		c := t[i]
		need[c]++
	}

	left, right := 0, 0
	valid := 0

	//记录最小覆盖字串的起始索引及长度
	start := 0
	length := math.MaxInt32
	for right < len(s) {
		c := s[right] //移入窗口的字符
		right++       //右移窗口

		//进行窗口内一系列数据更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		//判断左侧窗口是否要收紧
		for valid == len(need) {
			//更新最小覆盖字串
			if right-left < length {
				start = left
				length = right - left
			}
			//d是将移出窗口的字符
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

//func minWindow(s string, t string) string {
//	need, windows := map[byte]int{}, map[byte]int{}
//	for i := 0; i < len(t); i++ {
//		need[t[i]]++
//	}
//
//	left, right := 0, 0
//	valid := 0
//
//	start := 0
//	length := math.MaxInt
//
//	for right < len(s) {
//		c := s[right]
//		right++
//
//		if need[c] > 0 {
//			windows[c]++
//			if windows[c] == need[c] {
//				valid++
//			}
//		}
//
//		for len(need) == valid {
//			if right-left < length {
//				start = left
//				length = right - left
//			}
//
//			d := s[left]
//			left++
//			if need[d] > 0 {
//				if windows[d] == need[d] {
//					valid--
//				}
//				windows[d]--
//			}
//		}
//	}
//
//	if length == math.MaxInt {
//		return ""
//	}
//	return s[start : start+length]
//}

//func minWindow(s string, t string) string {
//	ori, cnt := map[byte]int{}, map[byte]int{}
//	for i := 0; i < len(t); i++ {
//		ori[t[i]]++
//	}
//
//	sLen := len(s)
//	len := math.MaxInt32
//	ansL, ansR := -1, -1
//
//	check := func() bool {
//		for k, v := range ori {
//			if cnt[k] < v {
//				return false
//			}
//		}
//		return true
//	}
//	for l, r := 0, 0; r < sLen; r++ {
//		if r < sLen && ori[s[r]] > 0 {
//			cnt[s[r]]++
//		}
//		for check() && l <= r {
//			if r-l+1 < len {
//				len = r - l + 1
//				ansL, ansR = l, l+len
//			}
//			if _, ok := ori[s[l]]; ok {
//				cnt[s[l]] -= 1
//			}
//			l++
//		}
//	}
//	if ansL == -1 {
//		return ""
//	}
//	return s[ansL:ansR]
//}
