package main

import (
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"os/exec"

	"github.com/mayuanucas/mygo/iot/config"
)

func init() {
	flag.BoolVar(&config.Version, "v", false, "显示版本信息")
	flag.StringVar(&config.OutputDir, "o", config.OUTPUTDIR, "指定输出文件夹")
	flag.StringVar(&config.InputDir, "i", "", "指定输入文件夹(该文件夹下为固件)")
	flag.StringVar(&config.InputDir2, "I", "", "指定输入文件夹(该文件夹的子目录下为固件)")
	flag.StringVar(&config.Command, "c", config.COMMAND, "指定解压分析固件脚本的路径")
}

func GetExtractScript() (string, error) {
	extractCommand := config.Command
	// 解压脚本不存在
	_, err := os.Stat(extractCommand)
	if err != nil && os.IsNotExist(err) {
		return "", err
	}

	return extractCommand, nil
}

func GetOutputDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	outputDir := config.OutputDir
	if strings.HasPrefix(outputDir, ".") {
		outputDir = cwd + string(os.PathSeparator) + outputDir[1:]
	}
	return outputDir, nil
}

func Task(extractCommand, firmwarePath, outputDir string) {
	cmd := exec.Command(extractCommand, firmwarePath, outputDir, ">/dev/null 2>&1")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("done->", firmwarePath)
}

func main() {
	flag.Parse()

	if config.Version {
		fmt.Printf("Firmware decompression analysis tool: version %s, developed by %s.\n", config.VERSION, config.AUTHOR)
		return
	}

	// 获取解压脚本路径
	extractCommand, err := GetExtractScript()
	if err != nil {
		fmt.Println("解压分析脚本不存在")
		return
	}

	if len(config.InputDir) < 1 && len(config.InputDir2) < 1 {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	} else if len(config.InputDir) >= 1 && len(config.InputDir2) >= 1 {
		fmt.Println("不支持 [-i -I] 同时使用.")
		return
	} else if len(config.InputDir) >= 1 {
		outputDir, err := GetOutputDir()
		if err != nil {
			fmt.Printf("保存路径无效 %s\n", config.OutputDir)
			return
		}

		files, _ := ioutil.ReadDir(config.InputDir)
		for _, file := range files {
			if file.IsDir() {
				continue
			}

			firmwarePath := config.InputDir + string(os.PathSeparator) + file.Name()
			Task(extractCommand, firmwarePath, outputDir)
		}
		fmt.Println("all done.")
	} else if len(config.InputDir2) >= 1 {
		outputDirRoot, err := GetOutputDir()
		if err != nil {
			fmt.Printf("保存路径无效 %s\n", config.OutputDir)
			return
		}

		dirs, _ := ioutil.ReadDir(config.InputDir2)
		for _, dir := range dirs {
			if !dir.IsDir() {
				continue
			}

			outputDir := outputDirRoot + string(os.PathSeparator) + dir.Name()
			// 解压分析该子目录下的所有固件
			childFiles, _ := ioutil.ReadDir(config.InputDir2 + string(os.PathSeparator) + dir.Name())
			for _, childFile := range childFiles {
				if childFile.IsDir() {
					continue
				}

				firmwarePath := config.InputDir2 + string(os.PathSeparator) +
					dir.Name() + string(os.PathSeparator) + childFile.Name()
				Task(extractCommand, firmwarePath, outputDir)
			}
		}
		fmt.Println("all done.")
	} else {
		return
	}
}
