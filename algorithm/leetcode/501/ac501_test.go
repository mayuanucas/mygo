package problem501

import (
	"testing"
	"fmt"
)

func Test_findMode(t *testing.T) {
	node21 := TreeNode{Val: 2, Left: nil, Right: nil}
	node2 := TreeNode{Val: 2, Left: &node21, Right: nil}
	node1 := TreeNode{Val: 1, Left: nil, Right: &node2}

	fmt.Println(findMode(&node1))
}
