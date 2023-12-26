package main

import "fmt"

/* 下一个更大元素 II */

//给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。
//
//数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1 。
//
//
//
//示例 1:
//
//输入: nums = [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；
//数字 2 找不到下一个更大的数；
//第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//示例 2:
//
//输入: nums = [1,2,3,4,3]
//输出: [2,3,4,-1,4]
//
//
//提示:
//
//1 <= nums.length <= 104
//-109 <= nums[i] <= 109

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}

// 循环数组可有两种方式，将nums两个一样的拼在一起，可作为循环数组计算，但可能会越界；使用i%length的方式=i
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stack := make([]int, 0)
	for i := 2*n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i%n] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) <= 0 {
			res[i%n] = -1
		} else {
			res[i%n] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%n])
	}
	return res
}
