package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* 二叉树的序列化与反序列化 */

func main() {
	fmt.Println(strings.Split("null,", ","))
	c := Constructor()
	//r := c.deserialize("1,2,null,null,3,4,null,null,5,null,null")
	r := c.deserialize("")
	fmt.Println(c.serialize(r))
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

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var str string
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			str += "null,"
			return
		}
		str += strconv.Itoa(node.Val) + ","
		preorder(node.Left)
		preorder(node.Right)
		return
	}
	preorder(root)
	return str[:len(str)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 0 {
		return nil
	}
	arr := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if arr[0] == "null" {
			arr = arr[1:]
			return nil
		}
		val, _ := strconv.Atoi(arr[0])
		arr = arr[1:]
		root := &TreeNode{Val: val}
		root.Left = build()
		root.Right = build()
		return root
	}
	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
