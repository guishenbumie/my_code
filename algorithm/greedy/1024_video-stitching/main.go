package main

import (
	"fmt"
	"sort"
)

/* 视频拼接 */

func main() {
	fmt.Println(videoStitching([][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10))
	fmt.Println(videoStitching([][]int{{0, 1}, {1, 2}}, 5))
	fmt.Println(videoStitching([][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}, 9))
}

func videoStitching(clips [][]int, time int) int {
	if time == 0 {
		return 0
	}
	sort.Slice(clips, func(i, j int) bool {
		clip1 := clips[i]
		clip2 := clips[j]
		if clip1[0] < clip2[0] {
			return true
		}
		if clip1[0] == clip2[0] {
			return clip1[1] > clip2[1]
		}
		return false
	})
	currEnd, nextEnd := 0, 0
	res := 0
	i := 0
	for i < len(clips) && clips[i][0] <= currEnd {
		for i < len(clips) && clips[i][0] <= currEnd {
			if clips[i][1] > nextEnd {
				nextEnd = clips[i][1]
			}
			i++
		}
		res++
		currEnd = nextEnd
		if currEnd >= time {
			return res
		}
	}
	return -1
}
