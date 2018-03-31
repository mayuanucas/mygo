package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mayuanucas/mygo/iot/config"
	"io/ioutil"
)

func init() {
	flag.BoolVar(&config.Version, "v", false, "显示版本信息")
	flag.StringVar(&config.OutputDir, "o", "./fdat", "指定输出文件夹")
	flag.StringVar(&config.InputDir, "i", "", "指定输入文件夹(该文件夹下为固件)")
	flag.StringVar(&config.InputDir2, "I", "", "指定输入文件夹(该文件夹的子目录下为固件)")
}

func main() {
	flag.Parse()

	if config.Version {
		fmt.Printf("Firmware decompression analysis tool: version %s, developed by %s.\n", config.VERSION, config.AUTHOR)
		return
	}

	if len(config.InputDir) < 1 && len(config.InputDir2) < 1 && len(config.OutputDir) < 1 {
		fmt.Printf("Usage of %s:", os.Args[0])
		flag.PrintDefaults()
		return
	}

	if len(config.InputDir) < 1 && len(config.InputDir2) < 1 {
		fmt.Println("请指定输入文件夹路径.")
		return
	} else if len(config.InputDir) >= 1 && len(config.InputDir2) >= 1 {
		fmt.Println("不支持 [-i -I] 同时使用.")
		return
	} else if len(config.InputDir) >= 1 {
		fmt.Println(config.OutputDir)
		files, _ := ioutil.ReadDir(config.InputDir)
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			fmt.Println(file.Name())
		}
	} else if len(config.InputDir2) >= 1 {
		dirs, _ := ioutil.ReadDir(config.InputDir2)
		for _, dir := range dirs {
			if !dir.IsDir() {
				continue
			}

			// 对每个目录建立对应解压后的子目录
			outputDir := config.OutputDir + string(os.PathSeparator) + dir.Name()
			e := os.Mkdir(outputDir, os.FileMode(0664))
			if e != nil {
				fmt.Println("创建目录失败:", outputDir)
				return
			}
			// 解压分析该子目录下的所有固件
			childFiles, _ := ioutil.ReadDir(config.InputDir2 + string(os.PathSeparator) + dir.Name())
			for _, childFile := range childFiles {
				if childFile.IsDir() {
					continue
				}
				fmt.Println(childFile.Name())
			}
		}
	} else {
		return
	}
}
