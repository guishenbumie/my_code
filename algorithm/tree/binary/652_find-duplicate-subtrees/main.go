package main

import (
	"strconv"
)

/* 寻找重复的子树 */

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	//res = list.New()
	//traverse(root)
	//arr := make([]*TreeNode, 0)
	//for e := res.Front(); e != nil; e = e.Next() {
	//	arr = append(arr, e.Value.(*TreeNode))
	//}
	//return arr
	mp := map[string]int{}
	strMp := map[string]*TreeNode{}
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return "|"
		}
		var str string
		str += "." + strconv.Itoa(root.Val)
		str += dfs(root.Left)
		str += dfs(root.Right)
		mp[str]++
		strMp[str] = root
		return str
	}
	dfs(root)
	res := []*TreeNode{}
	for i := range mp {
		if mp[i] > 1 {
			res = append(res, strMp[i])
		}
	}
	return res
}

//var memo = map[string]int{}
//var res = &list.List{}

//func traverse(root *TreeNode) string {
//	if root == nil {
//		return "#"
//	}
//	left := traverse(root.Left)
//	right := traverse(root.Right)
//	subTree := left + "," + right + "," + string(root.Val)
//
//	freq := memo[subTree]
//	if freq == 1 {
//		res.PushBack(root)
//	}
//	memo[subTree] = freq + 1
//	return subTree
//}
