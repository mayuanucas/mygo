package problem637

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if nil == root {
		return nil
	}

	ans := make([]float64, 0)
	// 模拟队列
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		n := len(queue)
		sum := 0.0
		for i := 0; i < n; i++ {
			node := queue[0]
			sum += float64(node.Val)
			if nil != node.Left {
				queue = append(queue, node.Left)
			}
			if nil != node.Right {
				queue = append(queue, node.Right)
			}
			// 从队列中删除头部元素，相当于头部元素出队
			queue = queue[1:]
		}
		ans = append(ans, sum/float64(n))
	}
	return ans
}
