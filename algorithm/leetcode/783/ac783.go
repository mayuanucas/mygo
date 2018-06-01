package problem783

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 因为给定的是二叉搜索树，故可以采用中序遍历 data[1]代表pre
func minDiffInBST(root *TreeNode) int {
	data := make([]int, 2)
	data[0] = math.MaxInt32
	data[1] = -1
	check(root, data)
	return data[0]
}

func check(node *TreeNode, nums []int) {
	if nil == node {
		return
	}
	check(node.Left, nums)
	if -1 != nums[1] {
		nums[0] = myMin(nums[0], node.Val-nums[1])
	}
	nums[1] = node.Val
	check(node.Right, nums)
}

func myMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
