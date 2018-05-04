package problem530

import (
	"testing"
	"fmt"
)

func Test_getMinimumDifference(t *testing.T) {
	node3 := TreeNode{Val: 3, Left: nil, Right: nil}
	node5 := TreeNode{Val: 5, Left: &node3, Right: nil}
	node1 := TreeNode{Val: 1, Left: nil, Right: &node5}

	fmt.Println(getMinimumDifference(&node1))

}
