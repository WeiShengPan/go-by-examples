package example8

import (
	"fmt"
	"testing"
)

// go routines
// 协程 在执行上来说是轻量级的线程
func TestGoRoutines(t *testing.T) {

	f("direct")

	go f("goroutine")

	// 匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

	// 当我们运行这个程序时，将首先看到阻塞式调用的输出，然后是两个 go 协程的交替输出。这种交替的情况表示 go 协程
	// 运行时是以异步的方式运行协程的。
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
