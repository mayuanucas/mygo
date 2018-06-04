package problem811

import (
	"strings"
	"strconv"
)

func subdomainVisits(cpdomains []string) []string {
	ans := make(map[string]int, 0)

	for _, item := range cpdomains {
		parts := strings.Split(item, " ")
		domain := parts[1]
		count, _ := strconv.Atoi(parts[0])

		ans[domain] += count
		idx := strings.LastIndex(domain, ".")
		for i := 0; i <= idx; i++ {
			if domain[i] == '.' {
				temp := domain[i+1:]
				ans[temp] += count
			}
		}
	}

	ansList := make([]string, 0)
	for k, v := range ans {
		ansList = append(ansList, strconv.Itoa(v)+" "+k)
	}
	return ansList
}
