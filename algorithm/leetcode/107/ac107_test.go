package problem107

import (
	"testing"
	"fmt"
)

func Test_levelOrderBottom(t *testing.T) {
	node15 := TreeNode{Val: 15, Left: nil, Right: nil}
	node7 := TreeNode{Val: 7, Left: nil, Right: nil}
	node9 := TreeNode{Val: 9, Left: nil, Right: nil}
	node20 := TreeNode{Val: 20, Left: &node15, Right: &node7}
	node3 := TreeNode{Val: 3, Left: &node9, Right: &node20}

	result := levelOrderBottom(&node3)
	fmt.Println(result)
}
