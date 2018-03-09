package problem30

import (
	"github.com/emirpasic/gods/stacks/arraystack"
	"errors"
	"fmt"
)

var (
	dataStack = arraystack.New()
	minStack = arraystack.New()
)

func push(number int) {
	dataStack.Push(number)

	if 0 == minStack.Size() {
		minStack.Push(number)
	}else {
		topNumber, _ := minStack.Peek()
		topN := topNumber.(int)

		if number < topN {
			minStack.Push(number)
		}else {
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

func Example() {
	push(3)
	push(4)
	push(2)
	push(6)
	push(5)

	pop()

	minNumber, ok := min()
	if ok != nil {
		fmt.Println(ok)
	}else {
		fmt.Println("当前栈中最小值为:", minNumber)
	}
}