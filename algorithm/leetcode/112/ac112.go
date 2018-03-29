package problem112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if nil == root {
		return false
	}

	// 判断叶子节点的值是否与指定值相等
	if nil == root.Left && nil == root.Right {
		return root.Val == sum
	} else {
		return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
	}
}
