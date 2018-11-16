package main

import (
	"flag"
	"fmt"
	"testing"
)

// 命令行标志
// 提供一个flag包，支持基本的命令行标志解析
func TestCommandLineFlags(t *testing.T) {

	// 基本的标记声明仅支持字符串、整数和布尔值选项。
	// 这里我们声明一个默认值为 `"foo"` 的字符串标志 `word`
	// 并带有一个简短的描述。这里的 `flag.String` 函数返回一个字
	// 符串指针（不是一个字符串值），在下面我们会看到是如何
	// 使用这个指针的。
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 执行命令行解析
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
