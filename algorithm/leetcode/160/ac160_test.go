package problem160

import (
	"testing"
	"fmt"
)

// 两个单链表的相交节点为 node4 该节点值为 4
// node11 -> node22
//                         -> node4 -> node5 -> node6
// node1 -> node2 -> node3
func Test_getIntersectionNode(t *testing.T) {
	node6 := ListNode{Val: 6, Next: nil}
	node5 := ListNode{Val: 5, Next: &node6}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	node22 := ListNode{Val: 22, Next: &node4}
	node11 := ListNode{Val: 11, Next: &node22}

	fmt.Println(getIntersectionNode(&node11, &node1).Val)
}
