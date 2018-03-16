package problem52

type ListNode struct {
	value int
	next  *ListNode
}

func findFirstCommonNode(head1, head2 *ListNode) *ListNode {
	length1 := getListLength(head1)
	length2 := getListLength(head2)
	var lengthDif int
	var listHeadLong, listHeadShort, firstCommonNode *ListNode
	if length1 >= length2 {
		lengthDif = length1 - length2
		listHeadLong = head1
		listHeadShort = head2
	} else {
		lengthDif = length2 - length1
		listHeadLong = head2
		listHeadShort = head1
	}

	for i := 0; i < lengthDif; i++ {
		listHeadLong = listHeadLong.next
	}
	for listHeadLong != listHeadShort && nil != listHeadLong && nil != listHeadShort {
		listHeadLong = listHeadLong.next
		listHeadShort = listHeadShort.next
	}
	firstCommonNode = listHeadLong
	return firstCommonNode
}

func getListLength(head *ListNode) int {
	var length int
	node := head
	for nil != node {
		length++
		node = node.next
	}
	return length
}
