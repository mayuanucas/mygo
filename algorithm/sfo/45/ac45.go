package problem45

import (
	"sort"
	"strconv"
)

type intSlice []int

func (p intSlice) Len() int {
	return len(p)
}

func (p intSlice) Less(i, j int) bool {
	return strconv.Itoa(p[i])+strconv.Itoa(p[j]) < strconv.Itoa(p[j])+strconv.Itoa(p[i])
}

func (p intSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func printMinNumber(numbers []int) string {
	if nil == numbers {
		return ""
	}

	sort.Sort(intSlice(numbers))
	var ans string
	for _, value := range numbers {
		ans += strconv.Itoa(value)
	}
	return ans
}
