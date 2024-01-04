package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树遍历框架
func traverse(root *TreeNode) {
	traverse(root.Left)
	traverse(root.Right)
}

// 二维矩阵遍历框架
func dfs(grid [][]int, i, j int, visited [][]bool) {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		//超出索引边界
		return
	}
	if visited[i][j] {
		//已经遍历过了
		return
	}
	//进入节点(i,j)
	visited[i][j] = true
	dfs(grid, i-1, j, visited) //上
	dfs(grid, i+1, j, visited) //下
	dfs(grid, i, j-1, visited) //左
	dfs(grid, i, j+1, visited) //右
}
