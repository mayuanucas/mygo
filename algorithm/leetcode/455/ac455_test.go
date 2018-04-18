package problem455

import (
	"testing"
	"fmt"
)

func Test_findContentChildren(t *testing.T) {
	g1 := []int{1, 2, 3}
	s1 := []int{1, 1}

	fmt.Println(findContentChildren(g1, s1))

	g2 := []int{1, 2}
	s2 := []int{1, 2, 3}
	fmt.Println(findContentChildren(g2, s2))
}
