package main

/* 两数之和 II - 输入有序数组 */

func main() {

}

// 使用二分查找
func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		left, right := i+1, len(numbers)-1 //数组是非递减数序，所以可能存在>1的相同的元素，为了避免找到一样的，子循环从+1开始找
		for left <= right {
			mid := left + (right-left)/2
			x := target - numbers[i]
			if numbers[mid] < x {
				left = mid + 1
			} else if numbers[mid] > x {
				right = mid - 1
			} else {
				return []int{i + 1, mid + 1}
			}
		}
	}
	return []int{-1, -1}
}
