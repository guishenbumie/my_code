package main

/* 寻找旋转排序数组中的最小值 II */

func main() {

}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else { //和153题不同的点，可能存在相同的值，right--再继续找
			right--
		}
	}
	return nums[left]
}
