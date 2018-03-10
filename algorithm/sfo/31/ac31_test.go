package problem31

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_isPopOrder(t *testing.T) {
	type argument struct {
		pushStack []int
		popStack  []int
	}
	type answer struct {
		value bool
	}
	type example struct {
		arg argument
		ans answer
	}

	examples := []example{
		{
			arg: argument{
				pushStack: []int{1, 2, 3, 4, 5},
				popStack:  []int{4, 5, 3, 2, 1},
			},
			ans: answer{
				value: true,
			},
		}, {
			arg: argument{
				pushStack: []int{1, 2, 3, 4, 5},
				popStack:  []int{4, 3, 5, 1, 2},
			},
			ans: answer{
				value: false,
			},
		},
	}

	ast := assert.New(t)
	for _, exam := range examples {
		ast.Equal(exam.ans.value, isPopOrder(exam.arg.pushStack, exam.arg.popStack),
			"压入栈:%v 弹出栈:%v", exam.arg.pushStack, exam.arg.popStack)
	}

}
