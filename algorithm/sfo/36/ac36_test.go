package problem36

import (
	"testing"
	"fmt"
)

func Test_conver(t *testing.T) {
	node16 := binaryTreeNode{value: 16, left: nil, right: nil}
	node12 := binaryTreeNode{value: 12, left: nil, right: nil}
	node8 := binaryTreeNode{value: 8, left: nil, right: nil}
	node4 := binaryTreeNode{value: 4, left: nil, right: nil}
	node14 := binaryTreeNode{value: 14, left: &node12, right: &node16}
	node6 := binaryTreeNode{value: 6, left: &node4, right: &node8}
	node10 := binaryTreeNode{value: 10, left: &node6, right: &node14}

	ptr := conver(&node10)
	fmt.Println(ptr.value)

}
