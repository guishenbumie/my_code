package main

import (
	"container/list"
	"strconv"
)

/* 滑动谜题 */

func main() {

}

func slidingPuzzle(board [][]int) int {
	m := 2
	n := 3
	var str string
	target := "123450"
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			str += strconv.Itoa(board[i][j])
		}
	}
	neighbor := [][]int{
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}

	q := list.New()
	visited := map[string]struct{}{}
	q.PushBack(str)
	visited[str] = struct{}{}

	var step int
	for q.Len() > 0 {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			curr := q.Remove(q.Front()).(string)
			if curr == target {
				return step
			}
			idx := 0
			for curr[idx] != '0' {
				idx++
			}
			for _, j := range neighbor[idx] {
				newStr := swap(curr, j, idx)
				if _, ok := visited[newStr]; !ok {
					q.PushBack(newStr)
					visited[newStr] = struct{}{}
				}
			}
		}
		step++
	}
	return -1
}

func swap(str string, i, j int) string {
	c := []byte(str)
	temp := c[i]
	c[i] = c[j]
	c[j] = temp
	return string(c)
}
