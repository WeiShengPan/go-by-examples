package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 互斥锁
func TestMutexes(t *testing.T) {

	var state = make(map[int]int)

	// mutex将同步对state的访问
	var mutex = &sync.Mutex{}

	// ops将记录我们对state的操作次数
	var ops int64 = 0

	// 运行100个go routine来重复读取state
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				// 为了确保这个 Go 协程不会再调度中饿死，我们
				// 在每次操作后明确的使用 `runtime.Gosched()`
				// 进行释放。这个释放一般是自动处理的，像例如
				// 每个通道操作后或者 `time.Sleep` 的阻塞调用后
				// 相似，但是在这个例子中我们需要手动的处理。
				runtime.Gosched()
			}
		}()
	}

	// 运行10个go routine来重复写入state
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	// 获取输出最终的操作计数
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	// 对state使用一个最终的锁，显示它是如何结束的
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
