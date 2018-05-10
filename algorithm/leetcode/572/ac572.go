package problem572

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if nil == s || nil == t {
		return s == t || nil == t
	}

	return isSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

// 从该节点起的 2 个子树是否完全相同
func isSame(t1, t2 *TreeNode) bool {
	if nil == t1 || nil == t2 {
		return t1 == t2
	}

	return t1.Val == t2.Val && isSame(t1.Left, t2.Left) && isSame(t1.Right, t2.Right)
}
