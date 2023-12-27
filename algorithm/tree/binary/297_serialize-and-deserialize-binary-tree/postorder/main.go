package main

import (
	"strconv"
	"strings"
)

func main() {

}

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
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			str += "null,"
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		str += strconv.Itoa(node.Val) + ","
	}
	postorder(root)
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
		last := arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		if last == "null" {
			return nil
		}
		val, _ := strconv.Atoi(last)
		root := &TreeNode{Val: val}
		root.Right = build()
		root.Left = build()
		return root
	}
	return build()
}
