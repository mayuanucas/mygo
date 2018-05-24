package problem687

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	ans := make([]int, 1)
	if nil != root {
		dfs(root, ans)
	}
	return ans[0]
}

func dfs(node *TreeNode, res []int) int {
	l := 0
	if nil != node.Left {
		l = dfs(node.Left, res)
	}
	r := 0
	if nil != node.Right {
		r = dfs(node.Right, res)
	}

	resl := 0
	if nil != node.Left && node.Left.Val == node.Val {
		resl = l + 1
	}
	resr := 0
	if nil != node.Right && node.Right.Val == node.Val {
		resr = r + 1
	}

	if res[0] < resl+resr {
		res[0] = resl + resr
	}

	if resl >= resr {
		return resl
	} else {
		return resr
	}
}
