package problem206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	for nil != head {
		nextNode := head.Next
		head.Next = newHead
		newHead = head
		head = nextNode
	}
	return newHead
}

func reverseList1(head *ListNode) *ListNode {
	return reverseCore(head, nil)
}

func reverseCore(head, newHead *ListNode) *ListNode {
	if nil == head {
		return newHead
	}

	nextNode := head.Next
	head.Next = newHead
	return reverseCore(nextNode, head)
}
