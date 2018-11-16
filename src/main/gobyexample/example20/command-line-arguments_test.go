package main

import (
	"fmt"
	"os"
	"testing"
)

// 命令行参数
func TestCommandLineArguments(t *testing.T) {

	// `os.Args` 提供原始命令行参数访问功能。注意，切片中
	// 的第一个参数是该程序的路径，并且 `os.Args[1:]`保存
	// 所有程序的的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

}
