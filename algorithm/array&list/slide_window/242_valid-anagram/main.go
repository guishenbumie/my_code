package main

import "fmt"

/* 有效的字母异位词 */

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//
//
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//示例 2:
//
//输入: s = "rat", t = "car"
//输出: false
//
//
//提示:
//
//1 <= s.length, t.length <= 5 * 104
//s 和 t 仅包含小写字母

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("ab", "a"))
}

// 两个map进行遍历比较方式，最简单的
//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	m1, m2 := map[byte]int{}, map[byte]int{}
//	for i := 0; i < len(s); i++ {
//		c := s[i]
//		m1[c]++
//	}
//	for i := 0; i < len(t); i++ {
//		c := t[i]
//		m2[c]++
//	}
//	if len(m1) != len(m2) {
//		return false
//	}
//	for k, v := range m1 {
//		if vv, ok := m2[k]; !ok {
//			return false
//		} else {
//			if vv != v {
//				return false
//			}
//		}
//	}
//	return true
//}

// 使用滑动窗口的方式
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		c := t[i]
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(t) {
			if valid == len(need) {
				return true
			}
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
	return false
}
