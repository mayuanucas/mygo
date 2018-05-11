package problem160

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if nil == headA || nil == headB {
		return nil
	}

	ptrA, ptrB := headA, headB;
	for ptrA != ptrB {
		if nil == ptrA {
			ptrA = headB
		} else {
			ptrA = ptrA.Next
		}

		if nil == ptrB {
			ptrB = headA
		} else {
			ptrB = ptrB.Next
		}
	}

	return ptrA
}
