package main

import (
	"fmt"
	"time"
	"errors"
)

func main() {
	var runes []rune
	for _, r := range "Hello,世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

func f2(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text, ":", i)
		time.Sleep(1 * time.Second)
	}
}

func f1(arg int) (int, error) {
	if 42 == arg {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

func modify(slice []int) {
	fmt.Printf("%p\n", &slice)
	slice[1] = 10
}

func test(a, b int, c float64) {
	fmt.Println(a + b + int(c))
}

func testMap() {
	m := make(map[string]int)
	m["a"] = 90
	m["b"] = 60
	fmt.Println(m)
	fmt.Println(len(m))

	value, status := m["a"]
	fmt.Println(value, status)

	value, status = m["aa"]
	fmt.Println(value, status)
}

func testArray() {
	//var a [5]int
	//fmt.Println(a[3])
	//fmt.Println(len(a))
	//fmt.Println(cap(a))

	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println(len(b))

	s := make([]string, 5)
	fmt.Println("emp:", s)
}

func IsEven(number int) bool {
	if number&1 == 0 {
		return true
	} else {
		return false
	}
}

func IsEven2(number int) bool {
	if number%2 == 0 {
		return true
	} else {
		return false
	}
}

func testEven() {
	LOOP := 100000
	num := 9

	start1 := time.Now()
	fmt.Println(start1)
	for i := 0; i < LOOP; i++ {
		IsEven(num)
	}
	end1 := time.Now()
	fmt.Println(end1)
	fmt.Println("位运算：", end1.Sub(start1))

	start2 := time.Now()
	fmt.Println(start2)
	for i := 0; i < LOOP; i++ {
		IsEven2(num)
	}
	end2 := time.Now()
	fmt.Println(end2)
	fmt.Println("求余运算：", end2.Sub(start2))
}
