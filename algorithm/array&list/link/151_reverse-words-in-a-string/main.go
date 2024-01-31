package main

import "strings"

/* 反转字符串中的单词 */

func main() {

}

func reverseWords(s string) string {
	s = strings.TrimSpace(s)

	left, right := len(s)-1, len(s)-1

	ans := []string{}
	for left >= 0 {
		for left >= 0 && s[left] != ' ' {
			left--
		}
		ans = append(ans, s[left+1:right+1])
		for left >= 0 && s[left] == ' ' {
			left--
		}
		right = left
	}
	return strings.Join(ans, " ")
}
