package problem59

import (
	"gopkg.in/oleiade/lane.v1"
)

func maxInWindows(num []int, size int) []int {
	maxInWindow := lane.NewDeque()

	if len(num) >= size && size >= 1 {
		index := lane.NewDeque()
		for i := 0; i < size; i++ {
			for !index.Empty() && num[i] >= num[index.Last().(int)] {
				index.Pop()
			}
			index.Append(i)
		}
		for i := size; i < len(num); i++ {
			maxInWindow.Append(num[index.First().(int)])
			for !index.Empty() && num[i] >= num[index.Last().(int)] {
				index.Pop()
			}
			if !index.Empty() && index.First().(int) <= i-size {
				index.Shift()
			}
			index.Append(i)
		}
		maxInWindow.Append(num[index.First().(int)])
	}
	var ans []int
	for !maxInWindow.Empty() {
		ans = append(ans, maxInWindow.Shift().(int))
	}
	return ans
}
