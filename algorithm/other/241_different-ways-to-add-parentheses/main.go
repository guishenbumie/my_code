package main

import (
	"fmt"
	"strconv"
)

/* 为运算表达式设计优先级 */

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}

// 分治算法
func diffWaysToCompute(expression string) []int {
	res := make([]int, 0)
	cs := []byte(expression)
	for i := 0; i < len(cs); i++ {
		c := cs[i]
		if c == '+' || c == '-' || c == '*' {
			//分
			left := diffWaysToCompute(string(cs[:i]))
			right := diffWaysToCompute(string(cs[i+1:]))
			//治
			for _, i := range left {
				for _, j := range right {
					if c == '+' {
						res = append(res, i+j)
					} else if c == '-' {
						res = append(res, i-j)
					} else if c == '*' {
						res = append(res, i*j)
					}
				}
			}
		}
	}
	if len(res) <= 0 { //如果计算结果为空，说明输入的字符串就是一个数字
		num, _ := strconv.Atoi(expression)
		res = append(res, num)
	}
	return res
}

//func diffWaysToCompute(input string) []int {
//	// 如果是数字，直接返回
//	if isDigit(input) {
//		tmp, _ := strconv.Atoi(input)
//		return []int{tmp}
//	}
//
//	// 空切片
//	var res []int
//	// 遍历字符串
//	for index, c := range input {
//		tmpC := string(c)
//		if tmpC == "+" || tmpC == "-" || tmpC == "*" {
//			// 如果是运算符，则计算左右两边的算式
//			left := diffWaysToCompute(input[:index])
//			right := diffWaysToCompute(input[index+1:])
//
//			for _, leftNum := range left {
//				for _, rightNum := range right {
//					var addNum int
//					if tmpC == "+" {
//						addNum = leftNum + rightNum
//					} else if tmpC == "-" {
//						addNum = leftNum - rightNum
//					} else {
//						addNum = leftNum * rightNum
//					}
//					res = append(res, addNum)
//				}
//			}
//		}
//	}
//
//	return res
//}
//
//// 判断是否为全数字
//func isDigit(input string) bool {
//	_, err := strconv.Atoi(input)
//	if err != nil {
//		return false
//	}
//	return true
//}
