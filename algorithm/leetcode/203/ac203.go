package problem203

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归写法 效率低
func removeElements(head *ListNode, val int) *ListNode {
	if nil == head {
		return nil
	}

	head.Next = removeElements(head.Next, val)

	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}

// 非递归 效率高
func removeElements2(head *ListNode, val int) *ListNode {
	for nil != head && val == head.Val {
		head = head.Next
	}
	currentPtr := head

	for nil != currentPtr && nil != currentPtr.Next {
		if currentPtr.Next.Val == val {
			currentPtr.Next = currentPtr.Next.Next
		} else {
			currentPtr = currentPtr.Next
		}
	}

	return head
}
