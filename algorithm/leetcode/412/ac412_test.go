package problem412

import (
	"testing"
	"fmt"
)

func Test_fizzBuzz(t *testing.T) {
	for _, v := range fizzBuzz(15) {
		fmt.Println(v)
	}
}
