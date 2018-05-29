package problem771

import (
	"testing"
	"fmt"
)

func Test_numJewelsInStones(t *testing.T) {
	j1 := "aA"
	s1 := "aAAbbbb"
	fmt.Println(numJewelsInStones(j1, s1))

	j2 := "z"
	s2 := "ZZ"
	fmt.Println(numJewelsInStones(j2, s2))
}
