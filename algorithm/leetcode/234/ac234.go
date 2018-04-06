package problem234

type ListNode struct {
	Val  int
	Next *ListNode
}

var h *ListNode

func isPalindrome(head *ListNode) bool {
	if nil == head {
		return true
	}

	if nil == h {
		h = head
	}
	flag := true
	if nil != head.Next {
		flag = flag && isPalindrome(head.Next)
	}

	flag = flag && (h.Val == head.Val)
	h = h.Next
	return flag
}

func isPalindrome2(head *ListNode) bool {
	nums := make([]int, 0, 128)
	for nil != head {
		nums = append(nums, head.Val)
		head = head.Next
	}

	for i, j := 0, len(nums)-1; i < j; {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}
	return true
}
