package problem530

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	numbers := []int{math.MaxInt32, -1}

	if nil == root {
		return numbers[0]
	}
	core(root, numbers)
	return numbers[0]
}

func core(root *TreeNode, nums []int) int {
	if nil == root {
		return nums[0]
	}

	core(root.Left, nums)

	if -1 != nums[1] {
		nums[0] = myMin(nums[0], root.Val-nums[1])
	}
	nums[1] = root.Val

	core(root.Right, nums)
	return nums[0]
}

func myMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
