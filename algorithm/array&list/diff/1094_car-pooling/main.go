package main

import "fmt"

/* 拼车 */

//车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）
//
//给定整数 capacity 和一个数组 trips ,  trip[i] = [numPassengersi, fromi, toi] 表示第 i 次旅行有 numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。
//
//当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。
//
//
//
//示例 1：
//
//输入：trips = [[2,1,5],[3,3,7]], capacity = 4
//输出：false
//示例 2：
//
//输入：trips = [[2,1,5],[3,3,7]], capacity = 5
//输出：true
//
//提示：
//
//1 <= trips.length <= 1000
//trips[i].length == 3
//1 <= numPassengersi <= 100
//0 <= fromi < toi <= 1000
//1 <= capacity <= 105

func main() {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
}

type Diff struct {
	diff []int
}

func Constructor(arr []int) Diff {
	if len(arr) <= 0 {
		return Diff{}
	}
	diff := make([]int, len(arr))
	diff[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		diff[i] = arr[i] - arr[i-1]
	}
	return Diff{diff: diff}
}

func (d *Diff) incr(i, j, v int) {
	d.diff[i] += v
	if j+1 < len(d.diff) {
		d.diff[j+1] -= v
	}
}

func (d *Diff) result() []int {
	arr := make([]int, len(d.diff))
	arr[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		arr[i] = arr[i-1] + d.diff[i]
	}
	return arr
}

func carPooling(trips [][]int, capacity int) bool {
	arr := make([]int, 1001)
	diff := Constructor(arr)
	for _, trip := range trips {
		diff.incr(trip[1], trip[2]-1, trip[0])
	}

	res := diff.result()
	for i := 0; i < len(res); i++ {
		if capacity < res[i] {
			return false
		}
	}
	return true
}
