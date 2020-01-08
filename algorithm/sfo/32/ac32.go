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
	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		e := l.Front()
		l.Remove(e)
		t := e.Value.(*binaryTreeNode)
		fmt.Print(t.value, " ")
		if t.left != nil {
			l.PushBack(t.left)
		}
		if t.right != nil {
			l.PushBack(t.right)
		}
	}
}

// 分行打印二叉树
func breadthFirstSearch2(root *binaryTreeNode) {
	if nil == root {
		return
	}
	l := list.New()
	l.PushBack(root)
	toBePrint := 1
	nextLevel := 0
	for l.Len() != 0 {
		e := l.Front()
		l.Remove(e)
		t := e.Value.(*binaryTreeNode)
		fmt.Printf("%d ",t.value)
		toBePrint--
		if t.left != nil {
			l.PushBack(t.left)
			nextLevel++
		}
		if t.right != nil {
			l.PushBack(t.right)
			nextLevel++
		}
		if toBePrint == 0 {
			fmt.Printf("\n")
			toBePrint = nextLevel
			nextLevel = 0
		}
	}

}
