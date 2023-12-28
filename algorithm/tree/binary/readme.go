package main

import (
	"container/list"
	"fmt"
)

func main() {

}

// 二叉树框架
type TreeNode struct {
	left  *TreeNode
	right *TreeNode
}

func Traverse(root *TreeNode) {
	//前序遍历
	Traverse(root.left)
	//中序遍历
	Traverse(root.right)
	//后序遍历
}

// 前序位置的代码只能从函数参数中获取父节点传递来的数据,
// 而后序位置的代码不仅可以获取参数数据,还可以获取到子树通过函数返回值专递回来的数据。

// 如果把根节点看做第1层,如何打印出每一个节点所在的层数?
func TraverseLv(root *TreeNode) {
	var f func(root *TreeNode, lv int)
	f = func(root *TreeNode, lv int) {
		if root == nil {
			return
		}
		fmt.Printf("节点%+v在第%d层\n", root, lv)
		f(root.left, lv+1)
		f(root.right, lv+1)
	}
	f(root, 1)
}

// 如何打印出每个节点的左右子树各有多少节点?
// 那么换句话说,一旦你发现题目和子树有关,那大概率要给函数设置合理的定义和返回值,在后序位置写代码了。
func Count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := Count(root.left)
	rightCount := Count(root.right)
	return leftCount + rightCount + 1
}

// 层级遍历二叉树
func LevelTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	//维护一个队列，迭代遍历的时候每次把根节点放进去
	l := list.New()
	l.PushBack(root)
	//从上到下遍历
	for l.Len() > 0 {
		//从左到右遍历
		sz := l.Len()
		for i := 0; i < sz; i++ {
			curr := l.Remove(l.Front()).(*TreeNode)
			//将下一层放到队列中
			if curr.left != nil {
				l.PushBack(curr.left)
			}
			if curr.right != nil {
				l.PushBack(curr.right)
			}
		}
	}
}
