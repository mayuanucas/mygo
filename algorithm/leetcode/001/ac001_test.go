package problem001

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type parameter struct {
	numberArray []int
	target      int
}

type answer struct {
	numberArray []int
}

type question struct {
	p parameter
	a answer
}

func TestOk(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: parameter{
				numberArray: []int{3, 2, 4},
				target:      6,
			},
			a: answer{
				numberArray: []int{1, 2},
			},
		}, {
			p: parameter{
				numberArray: []int{3, 2, 4},
				target:      8,
			},
			a: answer{
				numberArray: nil,
			},
		},
	}

	for _, q := range qs {
		ans, problem := q.a, q.p
		ast.Equal(ans.numberArray, twoSum(problem.numberArray, problem.target), "输入:%v", problem)
	}

}
