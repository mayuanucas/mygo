package problem55

type BinaryTreeNode struct {
	value       int
	left, right *BinaryTreeNode
}

func treeDepth(root *BinaryTreeNode) int {
	if nil == root {
		return 0
	}
	left := treeDepth(root.left)
	right := treeDepth(root.right)
	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}

func isBalanced(root *BinaryTreeNode, depth *int) bool {
	if nil == root {
		*depth = 0
		return true
	}
	var left, right int
	if isBalanced(root.left, &left) && isBalanced(root.right, &right) {
		diff := left - right
		if -1 <= diff && diff <= 1 {
			if left > right {
				*depth = 1 + left
			} else {
				*depth = 1 + right
			}
			return true
		}
	}
	return false
}

func isBalance(root *BinaryTreeNode) bool {
	var depth int
	return isBalanced(root, &depth)
}
