package problem111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	if 0 == leftDepth || 0 == rightDepth {
		return leftDepth + rightDepth + 1
	} else {
		if leftDepth < rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}
}
