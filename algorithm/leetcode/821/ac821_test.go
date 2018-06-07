package problem821

import (
	"testing"
	"fmt"
)

func Test_shortestToChar(t *testing.T) {
	S := "loveleetcode"
	C := byte('e')

	fmt.Println(shortestToChar(S, C))
}
