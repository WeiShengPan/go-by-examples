package main

import (
	"fmt"
	"strconv"
	"testing"
)

// 数字解析
// 内置的 strconv 包提供了数字解析功能
func TestNumberParsing(t *testing.T) {

	var p = fmt.Println

	// 使用 `ParseFloat` 解析浮点数，这里的 `64` 表示表示解
	// 析的数的位数。
	f, _ := strconv.ParseFloat("1.234", 64)
	p(f)

	// 在使用 `ParseInt` 解析整形数时，例子中的参数 `0` 表
	// 示自动推断字符串所表示的数字的进制。`64` 表示返回的
	// 整形数是以 64 位存储的。
	i, _ := strconv.ParseInt("123", 0, 64)
	p(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	p(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	p(u)

	// `Atoi` 是一个基础的 10 进制整型数转换函数。
	k, _ := strconv.Atoi("135")
	p(k)

	_, e := strconv.Atoi("wat")
	p(e)
}
