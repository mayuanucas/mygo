package problem30

import (
	"github.com/emirpasic/gods/stacks/arraystack"
	"errors"
)

var (
	dataStack = arraystack.New()
	minStack  = arraystack.New()
)

func push(number int) {
	dataStack.Push(number)

	if 0 == minStack.Size() {
		minStack.Push(number)
	} else {
		topNumber, _ := minStack.Peek()
		topN := topNumber.(int)

		if number < topN {
			minStack.Push(number)
		} else {
			minStack.Push(topN)
		}
	}
}

func pop() {
	if 0 < dataStack.Size() && 0 < minStack.Size() {
		dataStack.Pop()
		minStack.Pop()
	}
}

func min() (int, error) {
	if 0 < dataStack.Size() && 0 < minStack.Size() {
		topNumber, _ := minStack.Pop()
		topN := topNumber.(int)
		return topN, nil
	}
	return 0, errors.New("栈为空")
}
