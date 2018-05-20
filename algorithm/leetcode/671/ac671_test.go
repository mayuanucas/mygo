package problem671

import (
	"testing"
	"fmt"
)

func Test_findSecondMinimumValue(t *testing.T) {
	node35 := TreeNode{Val: 5}
	node37 := TreeNode{Val: 7}
	node22 := TreeNode{Val: 2}
	node25 := TreeNode{Val: 5, Left: &node35, Right: &node37}
	node12 := TreeNode{Val: 2, Left: &node22, Right: &node25}

	fmt.Println(findSecondMinimumValue(&node12))
}
