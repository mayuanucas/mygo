package main

import (
	"os/exec"
	"log"
	"flag"
	"fmt"
	"runtime"
	"io/ioutil"
	"os"
	"strings"
	"github.com/mayuanucas/mygo/lib/grpool"
	"github.com/mayuanucas/mygo/iot/i64/config"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func init() {
	flag.BoolVar(&config.Version, "v", false, "显示版本信息")
	flag.StringVar(&config.Interpreter, "C", config.INTERPRETER, "指定脚本的解释器")
	flag.StringVar(&config.Command, "c", config.COMMAND, "指定分析脚本的路径")
	flag.StringVar(&config.InputDir, "i", "", "指定输入文件夹(支持迭代)")
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
		fmt.Printf("生成I64文件多线程脚本运行工具: version %s, developed by %s.\n", config.VERSION, config.AUTHOR)
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

	handleBin(userCommand, config.InputDir, outputDir, pool)

	pool.WaitAll()
	log.Println("All tasks are completed.")
}

// 采用递归的方式处理多级目录下的文件
func handleBin(command, inputDir, outputDir string, pool *grpool.Pool) {
	files, _ := ioutil.ReadDir(inputDir)
	for _, file := range files {
		if file.IsDir() {
			handleBin(command,
				inputDir+string(os.PathSeparator)+file.Name(),
				outputDir+string(os.PathSeparator)+file.Name(),
				pool)
		} else {
			filePath := inputDir + string(os.PathSeparator) + file.Name()
			outputResultPath := outputDir + string(os.PathSeparator) + file.Name()

			// 指定保存目录不存在，则创建
			if !IsDir(outputDir) {
				os.MkdirAll(outputDir, 0775)
			}

			pool.Add(1)
			go Task(pool, config.Interpreter, command, filePath, outputResultPath)
		}
	}
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
	// 指定保存目录不存在，则创建
	if !IsDir(outputDir) {
		os.MkdirAll(outputDir, 0775)
	}

	return outputDir, nil
}

func Task(pool *grpool.Pool, programName string, arg ...string) {
	defer pool.Done()
	//忽略脚本的输出信息
	cmd := exec.Command(programName, arg...)
	err := cmd.Run()
	if err != nil {
		log.Println("error->", err)
	} else {
		log.Println("done->", arg[len(arg)-1])
	}
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	// os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
