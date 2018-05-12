package main

import (
	"time"
	"fmt"
)

func main()  {
	fmt.Println("现在时间是: ", time.Now())

	t := time.Now()
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Date())
	fmt.Println(t.Clock())

	fmt.Println(t.Format("19:21:01.123"))
}
