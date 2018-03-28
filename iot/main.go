package main

import (
	"flag"

	"github.com/mayuanucas/mygo/iot/config"
	"fmt"
	"os"
)

func init() {
	flag.BoolVar(&config.Version, "v", false, "显示版本信息")
	flag.StringVar(&config.OutputFile, "o", "", "指定输出文件名")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if config.Version {
		fmt.Printf("iot: version %s, A tool for iot, developed by %s.\n", config.VERSION, config.AUTHOR)
		return
	}

	if len(args) < 1 {
		fmt.Printf("Usage of %s: \n", os.Args[0])
		flag.PrintDefaults()
		return
	} else {
		fmt.Println(len(args), args)
	}
}
