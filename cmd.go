package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
	XjreOption  string //指定JRE目录位置
}

func parseCmd() *Cmd {
	// 一个叫做cmd的指针，指向了一个初始状态的Cmd结构体
	var cmd *Cmd = new(Cmd)
	flag.Usage = printUsage
	// 注册flag
	// 参数意义：用于存储解析后值的变量地址 ，用户在cmd行中使用的名称，默认值为false，，使用说明
	flag.BoolVar(&((*cmd).helpFlag), "h", false, "help")
	flag.BoolVar(&((*cmd).versionFlag), "v", false, "version")
	flag.StringVar(&((*cmd).cpOption), "cp", "", "classpath")
	flag.StringVar(&((*cmd).cpOption), "classpath", "", "classpath")
	flag.StringVar(&((*cmd).XjreOption), "Xjre", "", "Xjre path")
	flag.Parse()
	var args []string = flag.Args()
	if len(args) > 0 {
		(*cmd).class = args[0]
		// 从索引1开始到结尾的所有元素
		(*cmd).args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
