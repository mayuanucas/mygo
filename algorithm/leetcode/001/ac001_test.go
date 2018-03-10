package problem001

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type argument struct {
	numberArray []int
	target      int
}

type answer struct {
	numberArray []int
}

type example struct {
	arg argument
	ans answer
}

func TestOk(t *testing.T) {
	examples := []example{
		{
			arg: argument{
				numberArray: []int{3, 2, 4},
				target:      6,
			},
			ans: answer{
				numberArray: []int{1, 2},
			},
		}, {
			arg: argument{
				numberArray: []int{3, 2, 4},
				target:      8,
			},
			ans: answer{
				numberArray: nil,
			},
		},
	}

	ast := assert.New(t)
	for _, exam := range examples {
		arg, ans := exam.arg, exam.ans
		ast.Equal(ans.numberArray, twoSum(arg.numberArray, arg.target),
			"%v %v", arg.numberArray, arg.target)
	}

}
