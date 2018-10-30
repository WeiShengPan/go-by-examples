package main

import (
	"fmt"
	"os"
)

// 退出
func main() {

	// 当使用os.Exit时，defer不会执行，所以这里不会打印
	defer fmt.Println("!!!")

	os.Exit(3)
}
