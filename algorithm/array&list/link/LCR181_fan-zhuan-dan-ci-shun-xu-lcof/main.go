package main

import "strings"

/* 字符串中的单词反转 */

// 和160题一样
func main() {

}

func reverseMessage(message string) string {
	message = strings.TrimSpace(message)
	left, right := len(message)-1, len(message)-1
	ans := []string{}
	for left >= 0 {
		for left >= 0 && message[left] != ' ' {
			left--
		}
		ans = append(ans, message[left+1:right+1])
		for left >= 0 && message[left] == ' ' {
			left--
		}
		right = left
	}
	return strings.Join(ans, " ")
}
