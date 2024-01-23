package main

import "fmt"

/* 寻找旋转排序数组中的最小值 */

func main() {
	//fmt.Println(findMin([]int{3, 4, 5, 1, 2}))       //1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) //0
	fmt.Println(findMin([]int{11, 13, 15, 17}))      //11
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
