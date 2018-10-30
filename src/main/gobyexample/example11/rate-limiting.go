package example11

import (
	"fmt"
	"time"
)

// 速率限制
// 速率限制是一个重要的控制服务资源利用和质量的途径。go routine/channel/ticker优雅的支持了速率限制
func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter将每200ms接收一个值，这个是速率限制任务中的管理器
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		// 每次都阻塞limiter的接收，达到限制每200ms一次的效果
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们想临时进行速率限制，并且不影响整体的速率控制，我们可以通过通道缓冲来实现。
	// 这里的通道用于进行3次临时的脉冲型速率限制
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每200ms添加一个值到burstyLimiter中，直到达到3个的上限
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 现在模拟超过 5 个的接入请求。它们中刚开始的 3 个将由于受 burstyLimiter 的“脉冲”影响。
	// 第二批请求，我们直接连续处理了 3 次，这是由于这个“脉冲”
	// 速率控制，然后大约每 200ms 处理其余的 2 个。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
