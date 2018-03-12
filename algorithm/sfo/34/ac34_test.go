package problem34

import (
	"testing"
)

func Test_findPath(t *testing.T) {
	node4 := binaryTreeNode{value: 4, left: nil, right: nil}
	node7 := binaryTreeNode{value: 7, left: nil, right: nil}
	node5 := binaryTreeNode{value: 5, left: &node4, right: &node7}
	node12 := binaryTreeNode{value: 12, left: nil, right: nil}
	node10 := binaryTreeNode{value: 10, left: &node5, right: &node12}

	findPath(&node10, 22)
}
