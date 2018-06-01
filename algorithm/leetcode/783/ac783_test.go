package problem783

import (
	"testing"
	"fmt"
)

func Test_minDiffInBST(t *testing.T) {
	node44 := TreeNode{Val: 44}
	node50 := TreeNode{Val: 50, Left: &node44}
	node58 := TreeNode{Val: 58, Left: &node50}
	node34 := TreeNode{Val: 34, Right: &node58}
	node27 := TreeNode{Val: 27, Right: &node34}

	fmt.Println(minDiffInBST(&node27))
}
