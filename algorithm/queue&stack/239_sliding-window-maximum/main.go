package main

import (
	"fmt"
)

/* 滑动窗口最大值 */

func main() {
	//fmt.Println(maxSlidingWindow([]int{1}, 1))
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

//时间复杂度太高，n^2
//func maxSlidingWindow(nums []int, k int) []int {
//	res := make([]int, len(nums)-k+1)
//	for i := 0; i < len(res); i++ {
//		res[i] = math.MinInt
//	}
//
//	left, right := 0, k-1
//	for right < len(nums) {
//		for i := left; i <= right; i++ {
//			if res[left] < nums[i] {
//				res[left] = nums[i]
//			}
//		}
//		left++
//		right++
//	}
//
//	return res
//}

// 单调队列，这里用的单调递减队列，每次往队列中push一个元素，将前面比自己小的都删了，保持队列的单调性
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k+1)

	q := []int{}

	for i, x := range nums {
		//入队
		for len(q) > 0 && nums[q[len(q)-1]] < x {
			q = q[:len(q)-1] //维护队列单调递减
		}
		q = append(q, i)

		//出队
		if i-q[0] >= k { //队首已经离开窗口了，要去掉
			q = q[1:]
		}

		//维护答案
		if i >= k-1 {
			//由于队首到队尾单调递减，所以窗口最大值就是队首
			ans = append(ans, nums[q[0]])
		}
	}

	return ans
}
