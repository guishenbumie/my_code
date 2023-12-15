package main

import "fmt"

/* 字符串的排列 */

//给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
//换句话说，s1 的排列之一是 s2 的 子串 。
//
//
//
//示例 1：
//
//输入：s1 = "ab" s2 = "eidbaooo"
//输出：true
//解释：s2 包含 s1 的排列之一 ("ba").
//示例 2：
//
//输入：s1= "ab" s2 = "eidboaoo"
//输出：false
//
//
//提示：
//
//1 <= s1.length, s2.length <= 104
//s1 和 s2 仅包含小写字母

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}

func checkInclusion(s1 string, s2 string) bool {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s1); i++ {
		c := s1[i]
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[left]
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
