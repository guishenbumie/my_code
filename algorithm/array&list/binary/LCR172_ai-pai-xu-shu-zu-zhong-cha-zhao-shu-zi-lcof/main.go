package main

/* 统计目标成绩的出现次数 */

// 和34题一样
func main() {

}

func countTarget(scores []int, target int) int {
	left := leftBound(scores, target)
	if left < 0 {
		return 0
	}
	right := rightBound(scores, target)
	return right - left + 1
}

func leftBound(scores []int, target int) int {
	left, right := 0, len(scores)
	for left < right {
		mid := left + (right-left)/2
		if scores[mid] < target {
			left = mid + 1
		} else if scores[mid] > target {
			right = mid
		} else {
			right = mid
		}
	}
	return left
}

func rightBound(scores []int, target int) int {
	left, right := 0, len(scores)
	for left < right {
		mid := left + (right-left)/2
		if scores[mid] < target {
			left = mid + 1
		} else if scores[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
