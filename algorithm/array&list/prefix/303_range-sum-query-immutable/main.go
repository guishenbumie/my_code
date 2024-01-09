package main

import "fmt"

/* 区域和检索 - 数组不可变 */

// 给定一个整数数组  nums，处理以下类型的多个查询:
//
// 计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
// 实现 NumArray 类：
//
// NumArray(int[] nums) 使用数组 nums 初始化对象
// int sumRange(int i, int j) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，包含 left 和 right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )
//
// 示例 1：
//
// 输入：
// ["NumArray", "sumRange", "sumRange", "sumRange"]
// [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
// 输出：
// [null, 1, -1, -3]
//
// 解释：
// NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
// numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
// numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
// numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
// 0 -2 -2 1 -4 -2 -3

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

type NumArray struct {
	nums []int
}

// 前缀和的方式
func Constructor(nums []int) NumArray {
	preSums := make([]int, len(nums)+1)
	//索引=0的值就是0，方便计算
	for i := 1; i < len(preSums); i++ {
		preSums[i] = preSums[i-1] + nums[i-1]
	}
	return NumArray{nums: preSums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.nums[right+1] - this.nums[left]
}

/*
//普通的方式
func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left < 0 || right >= len(this.nums) {
		return 0
	}
	var n_sum int
	for i := left; i <= right; i++ {
		n_sum += this.nums[i]
	}
	return n_sum
}
*/
