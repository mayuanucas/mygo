package main

import "fmt"

/*
格式化输出语法

%d			十进制整数
%x %o %b	十六进制、八进制、二进制整数
%f %g %e	浮点数: 3.1415 3.141592653589793 3.141593e+00
%t			布尔: true 或 false
%c			字符（rune）(Unicode码点)
%s			字符串
%q			带双引号的字符串 "abc" 或带单引号的字符 'c'
%v			变量的自然形式
%T			变量的类型
%%			字面上的百分号标志（无操作数）
 */

 /*
 General

 	%v 以默认方式打印变量的值
    %T 打印变量的类型

Integer

    %+d 带符号的整型，fmt.Printf("%+d", -255)输出-255
    %q 打印单引号
    %o 不带零的八进制
    %#o 带零的八进制
    %x 小写的十六进制
    %X 大写的十六进制
    %#x 带0x的十六进制
    %U 打印Unicode字符
    %#U 打印带字符的Unicode
    %b 打印整型的二进制

Integer width

    %5d 表示该整型最大长度是5，下面这段代码

      fmt.Printf("|%5d|", 1)
      fmt.Printf("|%5d|", 1234567)

输出结果如下：

|    1|
|1234567|

    %-5d则相反，打印结果会自动左对齐
    %05d会在数字前面补零。

Float

    %f (=%.6f) 6位小数点
    %e (=%.6e) 6位小数点（科学计数法）
    %g 用最少的数字来表示
    %.3g 最多3位数字来表示
    %.3f 最多3位小数来表示

String

    %s 正常输出字符串
    %q 字符串带双引号，字符串中的引号带转义符
    %#q 字符串带反引号，如果字符串内有反引号，就用双引号代替
    %x 将字符串转换为小写的16进制格式
    %X 将字符串转换为大写的16进制格式
    % x 带空格的16进制格式

String Width (以5做例子）

    %5s 最小宽度为5
    %-5s 最小宽度为5（左对齐）
    %.5s 最大宽度为5
    %5.7s 最小宽度为5，最大宽度为7
    %-5.7s 最小宽度为5，最大宽度为7（左对齐）
    %5.3s 如果宽度大于3，则截断
    %05s 如果宽度小于5，就会在字符串前面补零

Struct

    %v 正常打印。比如：{sam {12345 67890}}
    %+v 带字段名称。比如：{name:sam phone:{mobile:12345 office:67890}
    %#v 用Go的语法打印。
    比如main.People{name:”sam”, phone:main.Phone{mobile:”12345”, office:”67890”}}

Boolean

    %t 打印true或false

Pointer

    %p 带0x的指针
    %#p 不带0x的指针

  */
func main() {
	learnScan()
}

func learnScan(){
	fmt.Println("请输入内容:")

	var n int
	fmt.Scanln(&n)

	numbers := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Println(n)
	fmt.Println(numbers)
}

func learnPrintf(){
	fmt.Printf("%05d\n", 1)
	fmt.Printf("|%5d|\n", 1234567)
	fmt.Printf("%-5d\n", 1)
	fmt.Printf("%+d\n", -3)
	fmt.Printf("二进制: %b\n", 3)
	fmt.Printf("八进制: %#o\n", 3)
	fmt.Printf("十六进制: %#x\n", 15)
	fmt.Printf("%f\n", 3.141592653589793)
	fmt.Printf("%g\n", 3.141592653589793)
	fmt.Printf("%.2e\n", 3.141592653589793)
	fmt.Printf("%s\n", "test\"")
	fmt.Printf("%q\n", "test\"")
	fmt.Printf("%#q\n", "test\"")
	fmt.Printf("%t\n", true)
	fmt.Printf("%v\n", false)
}
