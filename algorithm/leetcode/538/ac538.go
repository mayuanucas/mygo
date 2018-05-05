package problem538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if nil == root {
		return root
	}

	sums := []int{0}
	convertBSTCore(root, sums)
	return root
}

func convertBSTCore(root *TreeNode, sums []int) {
	if nil == root {
		return
	}

	convertBSTCore(root.Right, sums)

	root.Val += sums[0]
	sums[0] = root.Val

	convertBSTCore(root.Left, sums)
}
