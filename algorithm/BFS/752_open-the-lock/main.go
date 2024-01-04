package main

import (
	"container/list"
	"fmt"
)

/* 打开转盘锁 */

func main() {
	//fmt.Println(openLock([]string{"02"}, "01"))
	//fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"8888"}, "0009"))
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
}

func openLock(deadends []string, target string) int {
	deads := make(map[string]struct{})
	for _, v := range deadends {
		deads[v] = struct{}{}
	}
	visited := make(map[string]struct{})
	q := list.New()
	q.PushBack("0000")
	visited["0000"] = struct{}{}

	var step int
	for q.Len() > 0 {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			curr := q.Remove(q.Front()).(string)
			//判断是否到死亡终点
			if _, ok := deads[curr]; ok {
				continue
			}
			if curr == target {
				return step
			}
			for j := 0; j < 4; j++ {
				up := plus(curr, j)
				if _, ok := visited[up]; !ok {
					q.PushBack(up)
					visited[up] = struct{}{}
				}
				down := minus(curr, j)
				if _, ok := visited[down]; !ok {
					q.PushBack(down)
					visited[down] = struct{}{}
				}
			}
		}
		step++
	}
	return -1
}

// 密码上拨
func plus(s string, i int) string {
	c := []byte(s)
	if c[i] == '9' {
		c[i] = '0'
	} else {
		c[i] += 1
	}
	return string(c)
}

// 密码下拨
func minus(s string, i int) string {
	c := []byte(s)
	if c[i] == '0' {
		c[i] = '9'
	} else {
		c[i] -= 1
	}
	return string(c)
}
