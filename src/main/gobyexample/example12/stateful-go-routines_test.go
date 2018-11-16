package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

// go 状态协程
// 在前面的例子中，我们用互斥锁进行了明确的锁定来让共享的
// state 跨多个 Go 协程同步访问。另一个选择是使用内置的 Go
// 协程和通道的的同步特性来达到同样的效果。这个基于通道的方
// 法和 Go 通过通信以及    每个 Go 协程间通过通讯来共享内存，确
// 保每块数据有单独的 Go 协程所有的思路是一致的。
func TestStatefulGoRoutines(t *testing.T) {

	var ops int64

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动 100 个 Go 协程通过 `reads` 通道发起对 state 所有者
	// Go 协程的读取请求。每个读取请求需要构造一个 `readOp`，
	// 发送它到 `reads` 通道中，并通过给定的 `resp` 通道接收
	// 结果。
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 用相同的方法启动 10 个写操作。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

}

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}
