package main

/* 两数之和 II - 输入有序数组 */

// 和167题一样
func main() {

}

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		left, right := i+1, len(numbers)-1
		for left <= right {
			mid := left + (right-left)/2
			x := target - numbers[i]
			if numbers[mid] < x {
				left = mid + 1
			} else if numbers[mid] > x {
				right = mid - 1
			} else {
				return []int{i, mid}
			}
		}
	}
	return []int{-1, -1}
}
