package example10

import (
	"fmt"
	"testing"
	"time"
)

// 超时对于一个连接外部资源，或者其它一些需要花费执行时间的操作的程序而言是很重要的。
// 得益于channel和select，在 Go中实现超时操作是简洁而优雅的。
func TestChannelTimeouts(t *testing.T) {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
