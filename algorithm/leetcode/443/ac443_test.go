package problem443

import (
	"testing"
	"fmt"
)

func Test_compress(t *testing.T) {
	data := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(data))
}
