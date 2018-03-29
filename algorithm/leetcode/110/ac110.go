package problem110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}
	diff := depth(root.Left) - depth(root.Right)
	if diff > 1 || diff < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	left := depth(root.Left)
	right := depth(root.Right)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
