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
