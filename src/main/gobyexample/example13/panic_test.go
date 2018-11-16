package main

import (
	"os"
	"testing"
)

// Panic
// 以为这出乎意料的错误发生，通常用于表示程序正常运行中不应该出现或没有处理好的错误
// !!!非特殊情况，请不要在应用中使用panic
func TestPanic(t *testing.T) {

	//panic("a panic")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

}
