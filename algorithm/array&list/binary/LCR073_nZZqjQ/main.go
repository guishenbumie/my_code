package main

/* 爱吃香蕉的狒狒 */

// 和875题一样
func main() {

}

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1
	for _, p := range piles {
		if p > right {
			right = p
		}
	}
	for left < right {
		mid := left + (right-left)/2
		if calTime(piles, mid) < h {
			right = mid
		} else if calTime(piles, mid) > h {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func calTime(piles []int, speed int) int {
	time := 0
	for _, p := range piles {
		time += (p + speed - 1) / speed //向上取整
	}
	return time
}
