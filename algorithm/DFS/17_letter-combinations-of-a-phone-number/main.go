package main

import "fmt"

/* 电话号码的字母组合 */

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	mapping := []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	ans := []string{}
	if len(digits) == 0 {
		return ans
	}
	stk := []rune{}
	var dfs func(start int)
	dfs = func(start int) {
		if len(stk) == len(digits) {
			ans = append(ans, string(stk))
			return
		}
		for i := start; i < len(digits); i++ {
			d := digits[i] - '0'
			for _, c := range mapping[d] {
				stk = append(stk, c)
				dfs(i + 1)
				stk = stk[:len(stk)-1]
			}
		}
	}
	dfs(0)
	return ans
}
