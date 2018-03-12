package problem020

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	type argument struct {
		s string
	}
	type example struct {
		arg argument
		ans bool
	}
	examples := []example{
		{
			arg: argument{
				s: "{[()]}",
			},
			ans: true,
		},
		{
			arg: argument{
				s: "{{]]])",
			},
			ans: false,
		},
	}

	ast := assert.New(t)
	for _, exam := range examples {
		ans, arg := exam.ans, exam.arg
		ast.Equal(ans, isValid(arg.s),
			"%v %v", arg.s, ans)
	}
}
