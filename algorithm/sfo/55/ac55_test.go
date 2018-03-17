package problem55

import (
	"testing"
	"fmt"
)

func Test_treeDepth(t *testing.T) {
	node7 := BinaryTreeNode{value: 7, left: nil, right: nil}
	node6 := BinaryTreeNode{value: 6, left: nil, right: nil}
	node5 := BinaryTreeNode{value: 5, left: &node7, right: nil}
	node4 := BinaryTreeNode{value: 4, left: nil, right: nil}
	node3 := BinaryTreeNode{value: 3, left: nil, right: &node6}
	node2 := BinaryTreeNode{value: 2, left: &node4, right: &node5}
	node1 := BinaryTreeNode{value: 1, left: &node2, right: &node3}

	fmt.Println(treeDepth(&node1))

	fmt.Println(isBalance(&node1))

}
