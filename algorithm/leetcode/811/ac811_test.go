package problem811

import (
	"testing"
	"fmt"
)

func Test_subdomainVisits(t *testing.T) {
	test1 := []string{"9001 discuss.leetcode.com"}
	fmt.Println(subdomainVisits(test1))

	test2 := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	fmt.Println(subdomainVisits(test2))
}
