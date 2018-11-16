package main

import (
	"fmt"
	"os"
	"testing"
)

// 退出
func TestExit(t *testing.T) {

	// 当使用os.Exit时，defer不会执行，所以这里不会打印
	defer fmt.Println("!!!")

	os.Exit(3)
}
