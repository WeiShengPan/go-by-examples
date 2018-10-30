package example9

import "fmt"

// channel是连接多个 Go 协程的管道。你可以从一个 Go 协程将值发送到通道，然后在别的 Go 协程中接收。
func main() {

	// 使用 make(chan val-type) 创建一个新的通道, 后面的2表示最多允许缓存2个值
	messages := make(chan string, 2)

	// 使用 channel <- val 语法发送一个新的值到channel中，这里我们在一个新的go routine中发送 ping 到 messages中
	go func() {
		messages <- "ping"
	}()

	go func() {
		messages <- "pong"
	}()

	// 使用 <-channel 语法从channel中接收一个值，这里将接收发送的ping
	msg1 := <-messages
	msg2 := <-messages
	fmt.Println(msg1)
	fmt.Println(msg2)

	// 默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕，这个特性允许我们不适用任何其他的同步操作，在结尾等待ping
}
