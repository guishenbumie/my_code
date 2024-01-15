package main

import (
	"container/list"
	"fmt"
)

/* 基本计算器 II */

func main() {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate("(4+5)"))
	fmt.Println(calculate("4+5 *2"))
}

func isDigit(c int32) bool {
	return c >= '0' && c <= '9'
}

func calculate(s string) int {
	sl := list.New()
	for _, c := range s {
		if c == ' ' {
			continue
		}
		sl.PushBack(c)
	}

	var f func(sl *list.List) int
	f = func(sl *list.List) int {
		num := 0
		sign := '+'
		l := list.New()
		for sl.Len() > 0 {
			c := sl.Remove(sl.Front()).(int32)
			if isDigit(c) {
				num = 10*num + int(c-'0')
			}
			if c == '(' {
				num = f(sl)
			}
			if !isDigit(c) || sl.Len() == 0 {
				switch sign {
				case '+':
					l.PushBack(num)
				case '-':
					l.PushBack(-num)
				case '*':
					pre := l.Remove(l.Back()).(int)
					l.PushBack(pre * num)
				case '/':
					pre := l.Remove(l.Back()).(int)
					l.PushBack(pre / num)
				}
				num = 0
				sign = c
			}
			if c == ')' {
				break
			}
		}
		var res int
		for l.Len() > 0 {
			res += l.Remove(l.Back()).(int)
		}
		return res
	}

	return f(sl)
}
