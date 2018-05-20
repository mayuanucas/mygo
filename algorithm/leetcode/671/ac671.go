package problem671

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if nil == root {
		return -1
	}

	return findCore(root, root.Val)
}

func findCore(node *TreeNode, min int) int {
	if nil == node {
		return -1
	}
	if node.Val > min {
		return node.Val
	}

	leftMin := findCore(node.Left, min)
	rightMin := findCore(node.Right, min)
	if -1 == leftMin || -1 == rightMin {
		return myMax(leftMin, rightMin)
	} else {
		return myMin(leftMin, rightMin)
	}
}

func myMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func myMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
