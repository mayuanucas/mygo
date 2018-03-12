package problem32

import "testing"

func Test_breadthFirstSearch(t *testing.T) {
	var node6 = binaryTreeNode{value: 11, left: nil, right: nil}
	var node5 = binaryTreeNode{value: 9, left: nil, right: nil}
	var node4 = binaryTreeNode{value: 7, left: nil, right: nil}
	var node3 = binaryTreeNode{value: 5, left: nil, right: nil}
	var node2 = binaryTreeNode{value: 10, left: &node5, right: &node6}
	var node1 = binaryTreeNode{value: 6, left: &node3, right: &node4}
	var node0 = binaryTreeNode{value: 8, left: &node1, right: &node2}

	breadthFirstSearch(&node0)

	breadthFirstSearch2(&node0)
}
