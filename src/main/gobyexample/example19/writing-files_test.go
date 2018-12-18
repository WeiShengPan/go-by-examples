package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// 写文件
func TestWritingFiles(t *testing.T) {

	const path1 string = "d:/write1.txt"
	const path2 string = "d:/write2.txt"

	d1 := []byte("hello\r\ngo\r\n")
	err := ioutil.WriteFile(path1, d1, 0644)
	check2(err)

	f, err := os.Create(path2)
	check2(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check2(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// 将缓冲区信息写入磁盘
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	// 确保所有缓存的操作已写入底层写入器
	w.Flush()
}

func check2(e error) {
	if e != nil {
		panic(e)
	}
}
