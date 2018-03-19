package problem58

import (
	"fmt"
	"testing"
)

func Test_reverseSentence(t *testing.T) {
	text := "I am a student."

	fmt.Println(reverseSentence(text))
}

func Test_leftRotateString(t *testing.T) {
	text := "abcdefg"

	fmt.Println(leftRotateString(text, 2))
}
