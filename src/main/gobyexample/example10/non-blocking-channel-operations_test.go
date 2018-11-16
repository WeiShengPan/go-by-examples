package example10

import (
	"fmt"
	"testing"
)

// 非阻塞通道操作
// 常规的通过通道发送和接收数据是阻塞的。然而，我们可以使用带一个default子句的 select 来实现非阻塞的发送、接收，甚至是
// 非阻塞的多路 select。
func TestNonBlockingChannelOperations(t *testing.T) {

	messages := make(chan string)
	signals := make(chan bool)

	// 非阻塞接收，如果没有值，进入default分支
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 非阻塞发送，如果没有值，进入default分支
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 在default字句前使用多个case实现多路的非阻塞选择器
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
