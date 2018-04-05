package problem226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	return &TreeNode{Val: root.Val, Left: invertTree(root.Right), Right: invertTree(root.Left)}
}
