package problem836

// 两个矩形 A B 有重叠部分,当且仅当
// A.bottom < B.top && A.left < B.right and B.bottom < A.top && B.left < A.right
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if rec1[1] < rec2[3] && rec1[0] < rec2[2] && rec2[1] < rec1[3] && rec2[0] < rec1[2] {
		return true
	}
	return false
}
