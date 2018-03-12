package problem32

import (
	"gopkg.in/oleiade/lane.v1"
	"fmt"
)

type binaryTreeNode struct {
	value int
	left  *binaryTreeNode
	right *binaryTreeNode
}

func breadthFirstSearch(root *binaryTreeNode) {
	if nil == root {
		return
	}
	var deque = lane.NewDeque()
	deque.Append(root)

	for !deque.Empty() {
		ptr := deque.Shift().(*binaryTreeNode)
		fmt.Println(ptr.value)
		if ptr.left != nil {
			deque.Append(ptr.left)
		}
		if ptr.right != nil {
			deque.Append(ptr.right)
		}
	}
}

// 分行打印二叉树
func breadthFirstSearch2(root *binaryTreeNode) {
	if nil == root {
		return
	}
	var deque = lane.NewDeque()
	deque.Append(root)
	// 当前层中还没有打印的节点数
	currentLevel := 1
	// 下一层的节点数
	nextLevel := 0
	for !deque.Empty() {
		ptr := deque.First().(*binaryTreeNode)
		fmt.Print(ptr.value, " ")
		if ptr.left != nil {
			deque.Append(ptr.left)
			nextLevel++
		}
		if ptr.right != nil {
			deque.Append(ptr.right)
			nextLevel++
		}
		deque.Shift()
		currentLevel--
		if 0 == currentLevel {
			currentLevel = nextLevel
			nextLevel = 0
			fmt.Printf("\n")
		}
	}
}
