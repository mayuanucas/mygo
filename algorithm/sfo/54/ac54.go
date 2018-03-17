package problem54

type BinaryTreeNode struct {
	value       int
	left, right *BinaryTreeNode
}

func kthNode(root *BinaryTreeNode, k int) *BinaryTreeNode {
	if nil == root || k <= 0 {
		return nil
	}

	var target *BinaryTreeNode
	if root.left != nil {
		target = kthNode(root.left, k)
	}

	if target == nil {
		if 1 == k {
			target = root
		} else {
			k--
		}
	}

	if target == nil && root.right != nil {
		target = kthNode(root.right, k)
	}
	return target
}
