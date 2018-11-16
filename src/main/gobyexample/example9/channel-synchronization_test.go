package example9

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelSynchronization(t *testing.T) {

	done := make(chan bool, 1)
	go worker(done)

	// 程序将在接收到通道中worker发出的通知前一直阻塞, 这里如果把 <-done移除，程序甚至会在worker运行前结束
	<-done
	fmt.Println("exit")
}

// done 通道用于通知其他go routine这个函数已经工作完毕
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done")

	done <- true
}
