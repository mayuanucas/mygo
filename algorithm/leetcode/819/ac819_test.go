package problem819

import (
	"testing"
	"fmt"
)

func Test_mostCommonWord(t *testing.T) {
	para := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}

	fmt.Println(mostCommonWord(para, banned))
}
