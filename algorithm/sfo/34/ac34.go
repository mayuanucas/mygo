package problem34

import (
	"github.com/emirpasic/gods/stacks/arraystack"
	"fmt"
)

type binaryTreeNode struct {
	value int
	left  *binaryTreeNode
	right *binaryTreeNode
}

func findPath(root *binaryTreeNode, expectedSum int) {
	if nil == root {
		return
	}
	stack := arraystack.New()
	currentSum := 0
	findPathCore(root, stack, expectedSum, currentSum)
}

func findPathCore(root *binaryTreeNode, stack *arraystack.Stack, expectedSum, currentSum int) {
	currentSum += root.value
	stack.Push(root.value)

	// 如果是叶节点，并且路径节点值的和等于输入的期望值则打印路径
	isLeaf := root.left == nil && root.right == nil
	if isLeaf && currentSum == expectedSum {
		it := stack.Iterator()
		for it.End(); it.Prev(); {
			fmt.Print(it.Value(), " ")
		}
		fmt.Printf("\n")
	}

	// 如果不是叶节点，则遍历它的子节点
	if root.left != nil {
		findPathCore(root.left, stack, expectedSum, currentSum)
	}
	if root.right != nil {
		findPathCore(root.right, stack, expectedSum, currentSum)
	}
	stack.Pop()
}
