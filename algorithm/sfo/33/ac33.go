package problem33

// begin 开始字符的序号， end 结尾字符的序号
func verifySquenceOfBST(sequence []int, begin, end int) bool {
	if nil == sequence || begin >= end {
		return false
	}

	root := sequence[end]

	// 二叉搜索树中，左子树节点的值小于根节点的值
	i := begin
	for ; i < end; i++ {
		if sequence[i] > root {
			break
		}
	}
	// 二叉搜索树中，右子树节点的值大于根节点的值
	j := i
	for ; j < end; j++ {
		if sequence[j] < root {
			return false
		}
	}
	// 判断左子树是不是二叉搜索树
	left := true
	if 1 < i-begin {
		left = verifySquenceOfBST(sequence, begin, i-1)
	}
	// 判断右子树是不是二叉搜索树
	right := true
	if 1 < end-i {
		right = verifySquenceOfBST(sequence, i, end-1)
	}

	return left && right
}
