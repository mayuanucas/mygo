package main

import (
	"flag"
	"fmt"
	"runtime"
)

func usage() {
	fmt.Printf("iie 是一个工具-> %s %s\n", runtime.GOOS, runtime.GOARCH)
	inputDir := flag.String("i", "", "源文件夹路径")
	outputDir := flag.String("o", "./", "保存路径")
	flag.Parse()

	if "" != *inputDir {
		fmt.Println("输入路径:", *inputDir)
		fmt.Println("输出路径:", *outputDir)
	}else {
		fmt.Println("参数无效")
	}

}

func main() {
	usage()
}
