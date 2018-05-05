package problem501

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	ans := make([]int, 0)
	if nil == root {
		return ans
	}

	max, count, pre := 1, 0, root.Val
	traverse(root, &ans, &max, &count, &pre)
	return ans
}

func traverse(root *TreeNode, ans *[]int, max, count, pre *int) {
	if nil == root {
		return
	}

	traverse(root.Left, ans, max, count, pre)

	if root.Val == *pre {
		*count++
		if *count == *max {
			*ans = append(*ans, root.Val)
		} else if *count > *max {
			*max = *count
			*ans = (*ans)[:0]
			*ans = append(*ans, root.Val)
		}
	} else {
		*count = 1
		if *count == *max {
			*ans = append(*ans, root.Val)
		}
	}
	*pre = root.Val

	traverse(root.Right, ans, max, count, pre)

}
