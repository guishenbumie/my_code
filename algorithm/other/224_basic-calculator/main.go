package main

import (
	"container/list"
	"fmt"
)

/* 基本计算器 */

func main() {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate("(4+5)"))
	fmt.Println(calculate("4+5 *2"))
}

func isdigit(c int32) bool {
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
		l := list.New()
		num := 0
		sign := '+'
		for sl.Len() > 0 {
			c := sl.Remove(sl.Front()).(int32)
			if isdigit(c) {
				num = num*10 + int(c-'0')
			}
			if c == '(' {
				num = f(sl)
			}
			//if (!isdigit(c) && c != ' ') || sl.Len() == 0 {
			if !isdigit(c) || sl.Len() == 0 {
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
				sign = c
				num = 0
			}
			if c == ')' {
				//l.PushBack(num)
				break
			}
		}
		var res int
		for l.Len() > 0 {
			res += l.Remove(l.Back()).(int)
		}
		//fmt.Println(res)
		return res
	}
	return f(sl)
}
