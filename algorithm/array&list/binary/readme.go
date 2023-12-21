package main

func main() {

}

// 框架（在有序数组中搜索指定元素）
//func binarySearch(nums []int, target int) ... {
//	left, right := 0, ...
//	for ... {
//		//下面这行和(left+right)/2是一样的，只是为了防止left+right太大整数溢出
//		mid := left + (right-left)/2
//		if nums[mid] < target {
//			left = ...
//		} else if nums[mid] > target {
//			right = ...
//		} else {//也就是nums[mid] == target的场景
//			...
//		}
//	}
//	return ...
//}

//泛化，什么问题可以使用二分搜索算法技巧？
//首先,你要从题目中抽象出一个自变量x,一个关于x的函数f(x),以及一个目标值target。
//同时,x,f(x),target还要满足以下条件:
//1、f(x)必须是在x上的单调函数(单调增单调减都可以)。
//2、题目是让你计算满足约束条件f(x)== target时的x的值。

//// 函数f是关于自变量x的单调函数
//func f(x int) int {
//	// ...
//}
//
//// 主函数， 在f(x) == target 的约束下求x的最值
//func solution(nums []int, target int) int {
//	if len(nums) == 0 {
//		return -1
//	}
//	left := ... //自变量的最小值
//	right := ... + 1//自变量的最大值
//
//	for left < right {
//		mid := left + (right-left)/2
//		if f(mid) < target {
//			//怎么使f(x)大一点
//			...
//		} else if f(mid) > target {
//			//怎么使f(x)小一点
//			...
//		} else if f(mid) == target {
//			//题目是求左边届还是右边界
//			...
//		}
//	}
//
//	return left
//}
