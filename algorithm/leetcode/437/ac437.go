package problem437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if nil == root {
		return 0
	}
	return pathSumFrom(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func pathSumFrom(node *TreeNode, sum int) int {
	if nil == node {
		return 0
	}

	ans := 0
	if node.Val == sum {
		ans = 1
	}
	return ans + pathSumFrom(node.Left, sum-node.Val) + pathSumFrom(node.Right, sum-node.Val)
}
