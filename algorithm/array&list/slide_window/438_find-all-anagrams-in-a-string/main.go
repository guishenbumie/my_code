package main

import "fmt"

/* 找到字符串中所有字母异位词 */

//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
//异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
//
//
//
//示例 1:
//
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//示例 2:
//
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
//
//
//提示:
//
//1 <= s.length, p.length <= 3 * 104
//s 和 p 仅包含小写字母

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}

func findAnagrams(s string, p string) []int {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(p); i++ {
		c := p[i]
		need[c]++
	}

	arr := make([]int, 0)
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

		for right-left >= len(p) {
			if valid == len(need) {
				arr = append(arr, left)
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

	return arr
}
