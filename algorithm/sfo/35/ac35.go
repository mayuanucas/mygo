package problem35

type complexListNode struct {
	value   int
	next    *complexListNode
	sibling *complexListNode
}

func cloneNodes(head *complexListNode) {
	pNode := head
	for nil != pNode {
		pCloned := new(complexListNode)
		pCloned.value = pNode.value
		pCloned.next = pNode.next
		pNode.next = pCloned
		pNode = pCloned.next
	}

}

func connectSiblingNodes(head *complexListNode) {
	pNode := head
	for nil != pNode {
		pCloned := pNode.next
		if nil != pNode.sibling {
			pCloned.sibling = pNode.sibling.next
		}
		pNode = pCloned.next
	}
}

func reconnectNodes(head *complexListNode) *complexListNode {
	pNode := head
	var pClonedHead, pClonedNode *complexListNode
	if nil != pNode {
		pClonedHead = pNode.next
		pClonedNode = pClonedHead
		pNode.next = pClonedNode.next
		pNode = pNode.next
	}

	for nil != pNode {
		pClonedNode.next = pNode.next
		pClonedNode = pClonedNode.next
		pNode.next = pClonedNode.next
		pNode = pNode.next
	}
	return pClonedHead
}

func clone(pHead *complexListNode) *complexListNode {
	cloneNodes(pHead)
	connectSiblingNodes(pHead)
	return reconnectNodes(pHead)
}
