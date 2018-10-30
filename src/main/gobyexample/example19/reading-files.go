package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读文件
func main() {

	const path string = "src/main/gobyexample/example19/readme.txt"

	dat, err := ioutil.ReadFile(path)
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open(path)
	check(err)
	defer f.Close()

	// 读取部分 最多读取9个字节
	b1 := make([]byte, 9)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 使用Seek到一个文件中已知的位置并从这个位置开始读取（起始index是0，不是1）
	o2, err := f.Seek(3, 0)
	check(err)
	b2 := make([]byte, 6)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 3)
	n3, err := io.ReadAtLeast(f, b3, 3)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 没有内置的回转支持，使用Seek(0,0)实现
	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(9)
	check(err)
	fmt.Printf("9 bytes: %s\n", string(b4))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
