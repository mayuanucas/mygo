package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("请输入内容:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	text := strings.TrimSpace(input)

	fmt.Println("输入是:", text)

	fmt.Println("请输入数字:")
	input, _ = reader.ReadString('\n')
	number, _ := strconv.Atoi(strings.TrimSpace(input))
	fmt.Printf("%v,%T\n", number, number)

	fmt.Println("请输入多个数字:")
	input, _ = reader.ReadString('\n')
	temp := strings.TrimSpace(input)
	strs := strings.Split(temp, " ")
	fmt.Println(strs)
}
