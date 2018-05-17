package problem653

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if nil == root {
		return false
	}

	dict := make(map[int]bool, 0)
	return dfs(root, k, dict)
}

func dfs(node *TreeNode, k int, dict map[int]bool) bool {
	if nil == node {
		return false
	}
	if dict[k-node.Val] {
		return true
	}
	dict[node.Val] = true
	return dfs(node.Left, k, dict) || dfs(node.Right, k, dict)
}
