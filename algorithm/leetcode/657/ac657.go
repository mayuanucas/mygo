package problem657

import "strings"

type Position struct {
	x, y int
}

func judgeCircle(moves string) bool {
	pst := Position{x: 0, y: 0}

	for _, v := range moves {
		if 'R' == v {
			pst.x++
		} else if 'L' == v {
			pst.x--
		} else if 'U' == v {
			pst.y++
		} else if 'D' == v {
			pst.y--
		} else {
			continue
		}
	}

	if 0 == pst.x && pst.x == pst.y {
		return true
	}
	return false
}

func judgeCircle2(moves string) bool {
	r := strings.Count(moves, "R")
	l := strings.Count(moves, "L")
	u := strings.Count(moves, "U")
	d := strings.Count(moves, "D")

	if l == r && u == d {
		return true
	}
	return false
}
