package problem36

type binaryTreeNode struct {
	value       int
	left, right *binaryTreeNode
}

func conver(root *binaryTreeNode) *binaryTreeNode {
	var lastNodeInList *binaryTreeNode
	convertNode(root, &lastNodeInList)

	// lastNodeInList指向双向链表的尾节点，但是需要返回链表的头结点，所以进行下面操作。
	headOfList := lastNodeInList
	for nil != headOfList && nil != headOfList.left {
		headOfList = headOfList.left
	}
	return headOfList
}

func convertNode(node *binaryTreeNode, lastNodeInList **binaryTreeNode) {
	if nil == node {
		return
	}
	current := node
	if nil != current.left {
		convertNode(current.left, lastNodeInList)
	}
	current.left = *lastNodeInList
	if nil != *lastNodeInList {
		(*lastNodeInList).right = current
	}
	*lastNodeInList = current
	if nil != current.right {
		convertNode(current.right, lastNodeInList)
	}
}
