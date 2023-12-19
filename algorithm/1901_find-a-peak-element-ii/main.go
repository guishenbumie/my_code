package main

import "fmt"

/* 寻找峰值 II */

//一个 2D 网格中的 峰值 是指那些 严格大于 其相邻格子(上、下、左、右)的元素。
//
//给你一个 从 0 开始编号 的 m x n 矩阵 mat ，其中任意两个相邻格子的值都 不相同 。找出 任意一个 峰值 mat[i][j] 并 返回其位置 [i,j] 。
//
//你可以假设整个矩阵周边环绕着一圈值为 -1 的格子。
//
//要求必须写出时间复杂度为 O(m log(n)) 或 O(n log(m)) 的算法
//
//
//
//
//
//示例 1:
//-1 -1 -1 -1
//-1  1  4 -1
//-1  3  2 -1
//-1 -1 -1 -1
//
//
//输入: mat = [[1,4],[3,2]]
//输出: [0,1]
//解释: 3 和 4 都是峰值，所以[1,0]和[0,1]都是可接受的答案。
//示例 2:
//-1 -1 -1 -1 -1
//-1 10 20 15 -1
//-1 21 30 14 -1
//-1  7 16 32 -1
//-1 -1 -1 -1 -1
//
//
//输入: mat = [[10,20,15],[21,30,14],[7,16,32]]
//输出: [1,1]
//解释: 30 和 32 都是峰值，所以[1,1]和[2,2]都是可接受的答案。
//
//
//提示：
//
//m == mat.length
//n == mat[i].length
//1 <= m, n <= 500
//1 <= mat[i][j] <= 105
//任意两个相邻元素均不相等.

func main() {
	fmt.Println(findPeakGrid([][]int{{1, 4}, {3, 2}}))
	fmt.Println(findPeakGrid([][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}))
	//arr := [][]int{{0, 0, 0}, {1, 1, 1}}
	//fmt.Println(arr)
	//fmt.Println(arr[0])
}

// 找出一行里面的最大值
func max(row []int) int {
	idx := 0
	for i := 0; i < len(row); i++ {
		if row[i] > row[idx] {
			idx = i
		}
	}
	return idx
}

func findPeakGrid(mat [][]int) []int {
	rowNum := len(mat) //行数
	low, high := 0, rowNum-1
	for low <= high {
		mid := (low + high) / 2
		rowMaxIdx := max(mat[mid])
		if mid-1 >= 0 && mat[mid][rowMaxIdx] < mat[mid-1][rowMaxIdx] {
			high = mid - 1
			continue
		}
		if mid+1 < rowNum && mat[mid][rowMaxIdx] < mat[mid+1][rowMaxIdx] {
			low = mid + 1
			continue
		}
		return []int{mid, rowMaxIdx}
	}
	return nil
}
