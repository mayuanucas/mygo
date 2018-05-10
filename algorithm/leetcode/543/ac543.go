package problem543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的直径解题思路:
// 二叉树中的某个节点的左子树深度 与 该节点的右子树深度 之和为树种最大值，则该最大值即为二叉树的直径。
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	if nil == root {
		return ans
	}

	left := depthOfTree(root.Left, &ans)
	right := depthOfTree(root.Right, &ans)
	if left+right > ans {
		ans = left + right
	}

	return ans
}

func depthOfTree(node *TreeNode, diameter *int) int {
	if nil == node {
		return 0
	}

	leftDepth := depthOfTree(node.Left, diameter)
	rightDepth := depthOfTree(node.Right, diameter)
	temp := leftDepth + rightDepth
	if temp > *diameter {
		*diameter = temp
	}

	return myMax(leftDepth, rightDepth) + 1
}

func myMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
