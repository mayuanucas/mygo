package problem014

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	type argument struct {
		strs []string
	}
	type example struct {
		arg argument
		ans string
	}
	examples := []example{
		{
			arg: argument{
				strs: []string{"abcd", "abc", "ab"},
			},
			ans: "ab",
		},
	}

	ast := assert.New(t)
	for _, exam := range examples {
		ans, arg := exam.ans, exam.arg
		ast.Equal(ans, longestCommonPrefix(arg.strs),
			"%v %v", arg.strs, ans)
	}
}
