package problem506

import (
	"sort"
	"strconv"
)

type person struct {
	score int
	index int
}

type persons []person

// Len方法返回集合中的元素个数
func (p persons) Len() int {
	return len(p)
}

// Less方法报告索引i的元素是否比索引j的元素小
func (p persons) Less(i, j int) bool {
	// 此处实现从大到小排序 (大的数字在前，小的在后)
	return p[i].score > p[j].score
}

// Swap方法交换索引i和j的两个元素
func (p persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func findRelativeRanks(nums []int) []string {
	ans := make([]string, len(nums))

	ps := make(persons, len(nums))
	for i, v := range nums {
		ps[i].score = v
		ps[i].index = i
	}

	sort.Sort(ps)

	for i, v := range ps {
		switch i {
		case 0:
			ans[v.index] = "Gold Medal"
		case 1:
			ans[v.index] = "Silver Medal"
		case 2:
			ans[v.index] = "Bronze Medal"
		default:
			ans[v.index] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
