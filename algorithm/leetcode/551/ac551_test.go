package problem551

import (
	"testing"
	"fmt"
)

func Test_checkRecord(t *testing.T) {
	str1 := "PPALLP"
	fmt.Println(checkRecord(str1))

	str2 := "PPALLL"
	fmt.Println(checkRecord(str2))

}
