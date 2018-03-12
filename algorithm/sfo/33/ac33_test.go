package problem33

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_verifySquenceOfBST(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(true, verifySquenceOfBST([]int{5, 7, 6, 9, 11, 10, 8}, 0, 6),
		"是二叉搜索树的后序遍历序列")
	ast.Equal(false, verifySquenceOfBST([]int{7, 4, 6, 5}, 0, 3),
		"不是二叉搜索树的后序遍历序列")
}
