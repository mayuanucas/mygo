package problem009

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	type argument struct {
		x int
	}
	type example struct {
		arg argument
		ans bool
	}

	examples := []example{
		{
			arg: argument{
				x: 12321,
			},
			ans: true,
		},
		{
			arg: argument{
				x: 12344321,
			},
			ans: true,
		},
		{
			arg: argument{
				x: 98765,
			},
			ans: false,
		},
	}

	ast := assert.New(t)
	for _, exam := range examples {
		ans, arg := exam.ans, exam.arg
		ast.Equal(ans, isPalindrome(arg.x),
			"%v %v", arg.x, ans)
	}
}
