package main

import "fmt"

/* 有效的括号 */

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//每个右括号都有一个对应的相同类型的左括号。
//
//
//示例 1：
//
//输入：s = "()"
//输出：true
//示例 2：
//
//输入：s = "()[]{}"
//输出：true
//示例 3：
//
//输入：s = "(]"
//输出：false
//
//
//提示：
//
//1 <= s.length <= 104
//s 仅由括号 '()[]{}' 组成

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("(])"))
}

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	leftOf := func(v byte) byte {
		if v == '}' {
			return '{'
		}
		if v == ']' {
			return '['
		}
		return '('
	}
	stack := make([]byte, 0)
	for _, v := range s {
		vv := byte(v)
		if vv == '{' || vv == '[' || vv == '(' {
			stack = append(stack, vv)
		} else {
			if len(stack) == 0 { //只有右侧符号，肯定不符合
				return false
			}
			top := stack[len(stack)-1]
			if leftOf(vv) == top {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
