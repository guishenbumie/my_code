package main

import "fmt"

/* 二维区域和检索 - 矩阵不可变 */

//给定一个二维矩阵 matrix，以下类型的多个请求：
//
//计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
//实现 NumMatrix 类：
//
//NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
//int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。
//
//
//示例 1：
//
//3
//
//0 0
//0 3
//
//3 0
//5 6
//
//0 0 0
//0 3 3
//0 8 14
//
//3  0  1
//5  6  3
//1  2  0
//
//0  0  0  0
//0  3  3  4
//0  8  14 18
//0  9  17 21
//
//3  0  1  4  2
//5  6  3  2  1
//1  2  0  1  5
//
//0  0  0  0  0  0
//0  3  3  4  8  10
//0  8  14 18 24 27
//0  9  17 21 28 36

//
//3  0  1  4  2
//5  6  3  2  1
//1  2  0  1  5
//4  1  0  1  7
//1  0  3  0  5
//
//输入:
//["NumMatrix","sumRegion","sumRegion","sumRegion"]
//[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
//输出:
//[null, 8, 11, 12]
//
//解释:
//NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
//numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
//numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
//numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)

func main() {
	obj := Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	fmt.Println(obj.SumRegion(2, 1, 4, 3))
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
	fmt.Println(obj.SumRegion(1, 2, 2, 4))
}

type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])
	sums := make([][]int, m+1)
	sums[0] = make([]int, n+1)
	for i, row := range matrix {
		sums[i+1] = make([]int, n+1)
		for j, v := range row {
			sums[i+1][j+1] = sums[i+1][j] + sums[i][j+1] - sums[i][j] + v
		}
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
