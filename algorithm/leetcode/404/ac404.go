package problem404

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if nil == root {
		return 0
	}

	sum := 0
	if nil != root.Left && nil == root.Left.Left && nil == root.Left.Right {
		sum += root.Left.Val
	} else {
		sum += sumOfLeftLeaves(root.Left)
	}

	sum += sumOfLeftLeaves(root.Right)
	return sum
}
