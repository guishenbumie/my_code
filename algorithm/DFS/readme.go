package main

func main() {

}

//回溯算法，实际上就是一个决策树的遍历过程
//框架
//result = []
//def backtrack(路径,选择列表):
//	if 满足结束条件
//		result.add(路径)
//		return
//
//	for 选择 in 选择列表:
//		做选择
//		backtrack(路径,选择列表)
//		撤销选择
