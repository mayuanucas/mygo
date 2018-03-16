package problem50

import (
	"testing"
	"fmt"
)

func Test_firstNoRepeatingChar(t *testing.T) {
	str := "abaccdeff"
	ans := firstNoRepeatingChar(str)
	fmt.Printf("%c %v", ans, ans)
}
