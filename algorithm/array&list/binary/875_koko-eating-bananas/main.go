package main

import "fmt"

/* 爱吃香蕉的珂珂 */

//珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 h 小时后回来。
//
//珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。
//如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
//
//珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
//
//返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。
//
//
//
//示例 1：
//
//输入：piles = [3,6,7,11], h = 8
//输出：4
//示例 2：
//
//输入：piles = [30,11,23,4,20], h = 5
//输出：30
//示例 3：
//
//输入：piles = [30,11,23,4,20], h = 6
//输出：23
//
//
//提示：
//
//1 <= piles.length <= 104
//piles.length <= h <= 109
//1 <= piles[i] <= 109

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}

// 计算需要的时间，根据公式time=piles/speed
func calTime(piles []int, speed int) int {
	time := 0
	for _, p := range piles {
		time += (p + speed - 1) / speed //向上取整
	}
	return time
}

// 二分查找左边界问题
func minEatingSpeed(piles []int, h int) int {
	left := 1  //速度小，耗时长
	right := 1 //速度大，耗时短
	for _, p := range piles {
		if p > right {
			right = p
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if calTime(piles, mid) > h { //时间比目标长，速度太慢了，从mid右边再找
			left = mid + 1
		} else if calTime(piles, mid) < h { //时间比目标短，满足预期但速度快，从mid左边再找是否能找到更小的
			right = mid
		} else { //时间和目标相等，满足预期，从mid左边再找是否能找到更小的
			right = mid
		}
	}
	return left
}
