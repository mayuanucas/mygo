package main

import (
	"log"
	"flag"
	"os"
	"strings"
	"os/exec"
	"fmt"
	"io/ioutil"
	"runtime"

	"github.com/mayuanucas/mygo/lib/grpool"
	"github.com/mayuanucas/mygo/iot/feature8/config"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func init() {
	flag.BoolVar(&config.Version, "v", false, "显示版本信息")
	flag.StringVar(&config.Command, "c", config.COMMAND, "指定分析脚本的路径")
	flag.StringVar(&config.InputDir, "i", "", "指定输入文件夹(该文件夹下为二进制程序)")
	flag.StringVar(&config.OutputDir, "o", config.OUTPUTDIR, "指定输出文件夹")
	flag.Parse()
}

func main() {
	// 检查命令行参数
	if flag.NArg() > 0 || flag.NFlag() < 1 {
		flag.PrintDefaults()
		return
	}

	if config.Version {
		fmt.Printf("Feature extraction tool: version %s, developed by %s.\n", config.VERSION, config.AUTHOR)
		return
	}

	// 获取脚本路径
	userCommand, err := GetExtractScript()
	if err != nil {
		fmt.Println("分析脚本不存在.")
		return
	}
	// 检查输出目录是否合法
	outputDir, err := GetOutputDir()
	if err != nil {
		fmt.Printf("保存路径无效 %s\n", config.OutputDir)
		return
	}

	pool := grpool.New(runtime.NumCPU() + 1)
	files, _ := ioutil.ReadDir(config.InputDir)
	for _, file := range files {
		// 忽略该目录下的文件夹
		if file.IsDir() {
			continue
		}

		binFilePath := config.InputDir + string(os.PathSeparator) + file.Name()
		pool.Add(1)
		go Task(userCommand, binFilePath, outputDir, pool)
	}
	pool.WaitAll()
	log.Println("All tasks are completed.")
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

func Task(command, binFilePath, outputDir string, pool *grpool.Pool) {
	defer pool.Done()
	//忽略脚本的输出信息
	cmd := exec.Command(command, binFilePath, outputDir, ">/dev/null 2>&1")
	err := cmd.Run()
	if err != nil {
		log.Println("error->", err)
	}
	log.Println("done->", binFilePath)
}
