package problem563

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	ret := make([]int, 1)
	helper(root, ret)
	return ret[0]
}

func helper(root *TreeNode, ret []int) int {
	if nil == root {
		return 0
	}

	lSum := helper(root.Left, ret)
	rSum := helper(root.Right, ret)
	ret[0] += myAbs(lSum, rSum)

	return lSum + rSum + root.Val
}

func myAbs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
