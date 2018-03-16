package problem52

import (
	"testing"
	"fmt"
)

func Test_findFirstCommonNode(t *testing.T) {
	node7 := ListNode{value: 7, next: nil}
	node6 := ListNode{value: 6, next: &node7}
	node5 := ListNode{value: 5, next: &node6}
	node4 := ListNode{value: 4, next: &node5}
	node3 := ListNode{value: 3, next: &node6}
	node2 := ListNode{value: 2, next: &node3}
	node1 := ListNode{value: 1, next: &node2}

	fmt.Println(findFirstCommonNode(&node1, &node4).value)

}
