package example10

import (
	"fmt"
	"testing"
)

// 通道的关闭
// 关闭一个通道意味着不能再想这个通道发送值，可用于给通道的接收方传达工作已完成的信息
func TestClosingChannel(t *testing.T) {

	jobs := make(chan int, 5)
	done := make(chan bool)

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	// 关闭通道，通道中剩下的值仍然可以被接收
	close(jobs)
	fmt.Println("sent all jobs")

	// 使用j, more := <-jobs循环从jobs中接收数据，如果jobs已关闭，并且通道中所有的值都已经接受完毕，那么more值将为false。
	// 当完成所有任务时，通过done通道进行通知
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	<-done
}
