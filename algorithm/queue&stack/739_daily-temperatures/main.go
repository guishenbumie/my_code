package main

import "fmt"

/* 每日温度 */

//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//
//
//示例 1:
//
//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出: [1,1,4,2,1,1,0,0]
//示例 2:
//
//输入: temperatures = [30,40,50,60]
//输出: [1,1,1,0]
//示例 3:
//
//输入: temperatures = [30,60,90]
//输出: [1,1,0]
//
//
//提示：
//
//1 <= temperatures.length <= 105
//30 <= temperatures[i] <= 100

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			t := temperatures[top]
			if t <= temperatures[i] {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

		if len(stack) <= 0 {
			res[i] = 0
		} else {
			top := stack[len(stack)-1]
			res[i] = top - i
		}
		stack = append(stack, i)
	}
	return res
}

// 单调栈的框架，下面这个时间复杂度为O(n)，如果使用暴力法需要两个for循环，时间复杂度为O(n^2)
// 给你一个数组nums,请你返回一个等长的结果数组,结果数组中对应索引存储着下一个更大元素,如果没有更大的元素,就存-1。
// 比如说,输入一个数组nums [2,1,2,4,3],你返回数组[4,2,4,-1,-1]
func nextGreaterElement(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0)
	//倒着往栈里放
	for i := len(nums) - 1; i >= 0; i-- {
		//判断个子高矮
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			//矮个丢掉，因为被挡着了
			stack = stack[:len(stack)-1]
		}

		if len(stack) <= 0 { //后面没有比自己高的
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return res
}
