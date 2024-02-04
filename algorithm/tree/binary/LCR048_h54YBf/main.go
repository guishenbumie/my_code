package main

import (
	"strconv"
	"strings"
)

/* 二叉树的序列化与反序列化 */

// 和297题一样
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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	str := ""
	stk := []*TreeNode{}
	stk = append(stk, root)
	for len(stk) > 0 {
		sz := len(stk)
		for i := 0; i < sz; i++ {
			node := stk[i]
			if node == nil {
				str += "null,"
				continue
			}
			str += strconv.Itoa(node.Val)
			str += ","
			stk = append(stk, node.Left)
			stk = append(stk, node.Right)
		}
		stk = stk[sz:]
	}
	return str[:len(str)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, ",")
	rootVal, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: rootVal}
	stk := []*TreeNode{}
	stk = append(stk, root)
	for i := 0; i+1 < len(arr); i = i + 2 {
		parent := stk[0]
		stk = stk[1:]

		left := arr[i+1]
		if left != "null" {
			leftVal, _ := strconv.Atoi(left)
			parent.Left = &TreeNode{Val: leftVal}
			stk = append(stk, parent.Left)
		} else {
			parent.Left = nil
		}
		right := arr[i+2]
		if right != "null" {
			rightVal, _ := strconv.Atoi(right)
			parent.Right = &TreeNode{Val: rightVal}
			stk = append(stk, parent.Right)
		}
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
