package problem083

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	prePtr := head
	currentPtr := head.Next
	lastValue := prePtr.Val

	for nil != currentPtr {
		if currentPtr.Val != lastValue {
			lastValue = currentPtr.Val
			prePtr = currentPtr
			currentPtr = currentPtr.Next
		} else {
			prePtr.Next = currentPtr.Next
			currentPtr = currentPtr.Next
		}
	}
	return head
}
