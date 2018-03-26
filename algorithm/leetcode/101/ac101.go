package problem101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return nil == root || isSame(root.Left, root.Right)
}

func isSame(left, right *TreeNode) bool {
	if nil == left || nil == right {
		return left == right
	} else if left.Val != right.Val {
		return false
	} else {
		return isSame(left.Left, right.Right) && isSame(left.Right, right.Left)
	}
}
