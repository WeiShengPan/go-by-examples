package example6

import (
	"fmt"
	"testing"
)

// 闭包
// 这个 `intSeq` 函数返回另一个在 `intSeq` 函数体内定义的
// 匿名函数。这个返回的函数使用闭包的方式 _隐藏_ 变量 `i`。
func TestClosures(t *testing.T) {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
