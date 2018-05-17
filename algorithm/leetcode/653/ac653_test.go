package problem653

import (
	"testing"
	"fmt"
)

func Test_findTarget(t *testing.T) {
	node2 := TreeNode{Val: 2, Left: nil, Right: nil}
	node4 := TreeNode{Val: 4, Left: nil, Right: nil}
	node7 := TreeNode{Val: 7, Left: nil, Right: nil}
	node3 := TreeNode{Val: 3, Left: &node2, Right: &node4}
	node6 := TreeNode{Val: 2, Left: nil, Right: &node7}
	node5 := TreeNode{Val: 5, Left: &node3, Right: &node6}

	fmt.Println(findTarget(&node5, 9))
}
