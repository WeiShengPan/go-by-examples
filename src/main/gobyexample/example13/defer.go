package main

import (
	"fmt"
	"os"
)

// Defer
// 用于确保一个函数调用在程序执行结束前执行，类似java finally
func main() {
	f := createFile("src/main/gobyexample/example13/defer.txt")
	// 即使写在前面，这里的closeFile也会在main结束前执行，也就是writeFile结束后
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "this is data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
