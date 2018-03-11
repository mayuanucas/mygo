package problem013

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	type argument struct {
		str string
	}
	type example struct {
		arg argument
		ans int
	}

	examples := []example{
		{
			arg: argument{
				str: "VIII",
			},
			ans: 8,
		},
		{
			arg: argument{
				str: "IV",
			},
			ans: 4,
		},
	}

	ast := assert.New(t)
	for _, exam := range examples {
		ans, arg := exam.ans, exam.arg
		ast.Equal(ans, romanToInt(arg.str),
			"%v %v", arg.str, ans)
	}
}
