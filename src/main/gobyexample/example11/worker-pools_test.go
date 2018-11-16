package example11

import (
	"fmt"
	"testing"
	"time"
)

// 使用go routine和channel实现一个pool
func TestWorkerPools(t *testing.T) {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动3个worker，初始是阻塞的，因为没有传递任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送9个job，然后close来表示这是所有任务了
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后收集所有任务的返回值
	for a := 1; a <= 9; a++ {
		<-results
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		// sleep 1s模拟耗时操作
		time.Sleep(time.Second)
		results <- j * 2
	}
}
