package problem104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}
