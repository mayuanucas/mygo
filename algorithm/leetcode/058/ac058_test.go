package problem058

import (
	"testing"
	"fmt"
)

func Test_lengthOfLastWord(t *testing.T) {
	text1 := "Hello World"
	text2 := "a "
	fmt.Println(lengthOfLastWord(text1))
	fmt.Println(lengthOfLastWord(text2))

	fmt.Println(lengthOfLastWord2(text1))
	fmt.Println(lengthOfLastWord2(text2))

	text3 := "你好golang"
	fmt.Println(len(text3))
	fmt.Printf("%v %q %T\n", text3[0], text3[0], text3[0])
	fmt.Printf("%v %q %T\n", text3[6], text3[6], text3[6])

	text4 := '中'
	text5 := 'g'
	fmt.Printf("%v %T\n", text4, text4)
	fmt.Printf("%q %T\n", text5, text5)

}
