package example9

import (
	"fmt"
	"testing"
)

// 当使用通道作为函数的参数时，你可以指定这个通道是不是只用来发送或者接收值。这个特性提升了程序的类型安全性。
func TestChannelDirections(t *testing.T) {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

// 定义了一个只允许发送数据的channel，尝试通过此channel接收数据将得到一个编译时错误
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 允许通过pings来接收数据，pongs来发送数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
