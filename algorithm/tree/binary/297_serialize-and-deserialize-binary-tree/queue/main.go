package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// 层级遍历，是最标准的二叉树遍历的方式
func main() {
	//fmt.Println(strings.Split("null,", ","))
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	c := Constructor()
	str := c.serialize(root)
	fmt.Println(str)
	rr := c.deserialize(str)
	fmt.Println(c.serialize(rr))
	//r := c.deserialize("1,2,3,null,null,4,5")
	//l := list.New()
	//l.PushBack(&TreeNode{Val: 1})
	//fmt.Println(l.Len())
	//e := l.Front()
	//fmt.Println(e)
	//fmt.Println(l.Remove(e))
	//n := e.Value.(*TreeNode)
	//fmt.Println(n.Val)
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
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		curr := l.Front()
		l.Remove(curr)
		node := curr.Value.(*TreeNode)
		if node == nil {
			str += "null,"
			continue
		}
		str += strconv.Itoa(node.Val)
		str += ","
		l.PushBack(node.Left)
		l.PushBack(node.Right)
	}
	return str[:len(str)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 0 {
		return nil
	}
	arr := strings.Split(data, ",")
	rootVal, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: rootVal}
	l := list.New()
	l.PushBack(root)
	for i := 0; i+1 < len(arr); i = i + 2 {
		e := l.Front()
		l.Remove(e)
		parent := e.Value.(*TreeNode)

		left := arr[i+1]
		if left != "null" {
			leftVal, _ := strconv.Atoi(left)
			parent.Left = &TreeNode{Val: leftVal}
			l.PushBack(parent.Left)
		} else {
			parent.Left = nil
		}

		right := arr[i+2]
		if right != "null" {
			rightVal, _ := strconv.Atoi(right)
			parent.Right = &TreeNode{Val: rightVal}
			l.PushBack(parent.Right)
		} else {
			parent.Right = nil
		}
	}
	return root
}
