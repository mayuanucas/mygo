package problem007

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type argument struct {
	number int
}

type answer struct {
	revertNumber int
}

type example struct {
	arg argument
	ans answer
}

func Test_reverse(t *testing.T) {
	examples := []example{
		{
			arg: argument{
				number: 1230,
			},
			ans: answer{
				revertNumber: 321,
			},
		},
		{
			arg: argument{
				number: -123,
			},
			ans: answer{
				revertNumber: -321,
			},
		},
	}

	ast := assert.New(t)
	for _, exam := range examples {
		ans, arg := exam.ans, exam.arg
		ast.Equal(ans.revertNumber, reverse(arg.number),
			"%v %v", arg.number, ans.revertNumber)
	}
}
