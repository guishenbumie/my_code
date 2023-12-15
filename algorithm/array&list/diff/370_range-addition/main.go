package main

import "fmt"

//假设你有一个长度为n的数组,初始情况下所有的数字均为0,你将会被给出k个更新的操作。
//其中,每个操作会被表示为一个三元组:[startindex,endindex,inc],你需要将子数
//组A[startindex...endindex](包括startindex和endindex)增加inc。
//请你返回 k次操作后的数组。
//
//示例:
//输入: length = 5, updates = [[1,3,2],[2,4,3],[0,2,-2]]
//输出:[-2,0,3,5,3]
//解释:
//初始状态:
//[0,0,0,0,0]
//进行了操作[1,3,2]后的状态:
//[0,2,2,2,0]
//进行了操作[2,4,3]后的状态:
//[0,2,5,5,3]
//进行了操作[0,2,-2]后的状态:
//[-2,0,3,5,3]

func main() {
	fmt.Println(getModifiedArray(5, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}))
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

func getModifiedArray(length int, updates [][]int) []int {
	arr := make([]int, length)
	diff := Constructor(arr)
	for _, u := range updates {
		diff.incr(u[0], u[1], u[2])
		//fmt.Println(diff.result())
	}
	return diff.result()
}
