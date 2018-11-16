package example9

import (
	"fmt"
	"testing"
	"time"
)

// 通道选择器可以同时等待多个通道操作。Go 协程和通道以及选择器的结合是 Go 的一个强大特性。
func TestChannelSelect(t *testing.T) {

	c1 := make(chan string)
	c2 := make(chan string)

	// 各个通道将在若干时间后接收一个值，用于模拟并行的go routine中阻塞的rpc操作
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "two"
	}()
	go func() {
		time.Sleep(time.Second * 4)
		c2 <- "one"
	}()

	// 使用select来同时等待这两个值。这里因为只有2个go routine，如果i大于2，会提示 fatal error: all goroutines are asleep - deadlock!
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received ", msg1)
		case msg2 := <-c2:
			fmt.Println("received ", msg2)
		}
	}

}
