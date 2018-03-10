package problem31

import "github.com/emirpasic/gods/stacks/arraystack"

func isPopOrder(pushStack, popStack []int) bool {
	length := len(pushStack)
	if 0 == length || length != len(popStack) {
		return false
	}

	stack := arraystack.New()
	nextPush, nextPop := 0, 0

	for nextPop < length {
		for topNumber, _ := stack.Peek(); nil == topNumber || topNumber != popStack[nextPop]; {
			if nextPush == length {
				break
			}
			stack.Push(pushStack[nextPush])
			nextPush++

			topNumber, _ = stack.Peek()
		}
		if topNumber, _ := stack.Peek(); topNumber != popStack[nextPop] {
			break
		}

		stack.Pop()
		nextPop++
	}

	if stack.Empty() && nextPop == length {
		return true
	}
	return false
}

func Ma(){
	return
}
