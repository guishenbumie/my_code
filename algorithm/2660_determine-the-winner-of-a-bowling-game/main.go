package main

import "fmt"

/* 保龄球游戏的获胜者 */

func main() {
	fmt.Println(isWinner([]int{4, 10, 7, 9}, []int{6, 5, 2, 3}))
	fmt.Println(isWinner([]int{3, 5, 7, 6}, []int{8, 10, 10, 2}))
	fmt.Println(isWinner([]int{2, 3}, []int{4, 1}))
	fmt.Println(isWinner([]int{5, 6, 1, 10}, []int{5, 1, 10, 5}))
}

func isWinner(player1 []int, player2 []int) int {
	sum := func(nums []int) int {
		c := 0
		arr := make([]bool, len(nums)+2)
		for i := 0; i < len(nums); i++ {
			c += nums[i]
			if nums[i] == 10 {
				arr[i+1] = true
				arr[i+2] = true
			}
			if arr[i] {
				c += nums[i]
			}
		}
		return c
	}
	p1 := sum(player1)
	p2 := sum(player2)
	if p1 > p2 {
		return 1
	} else if p1 < p2 {
		return 2
	}
	return 0
}
