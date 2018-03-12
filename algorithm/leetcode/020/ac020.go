package problem020

import "github.com/emirpasic/gods/stacks/arraystack"

func isValid(s string) bool {
	stack := arraystack.New()
	for _, c := range s {
		if c == '(' {
			stack.Push(')')
		} else if c == '[' {
			stack.Push(']')
		} else if c == '{' {
			stack.Push('}')
		} else {
			if stack.Empty() {
				return false
			}
			tempC, _ := stack.Pop()
			if tempC != c {
				return false
			}
		}
	}
	return stack.Empty()
}
