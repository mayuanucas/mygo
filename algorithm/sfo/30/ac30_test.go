package problem30

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	ast := assert.New(t)

	push(3)
	push(4)
	push(2)
	push(6)
	push(5)
	pop()

	minNumber, ok := min()
	if nil != ok {
		ast.Error(ok)
	}
	ast.Equal(2, minNumber, "栈中最小数")
}
