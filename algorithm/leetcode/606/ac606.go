package problem606

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if nil == t {
		return ""
	}

	ans := strconv.Itoa(t.Val)
	strLeft := tree2str(t.Left)
	strRight := tree2str(t.Right)

	if "" == strLeft && "" == strRight {
		return ans
	}
	if "" == strLeft {
		return ans + "()" + "(" + strRight + ")"
	}
	if "" == strRight {
		return ans + "(" + strLeft + ")"
	}
	return ans + "(" + strLeft + ")" + "(" + strRight + ")"
}
