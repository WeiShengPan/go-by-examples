package example11

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// 原子计数器
// Go中最主要的状态管理方式是通过通道间的沟通来完成的，在worker-pool.go中有遇到过。这里还可以用sync/atomic包在多个go routine
// 中进行原子计数
func main() {

	// 使用一个无符号整形数表示这个计数器
	var ops uint64 = 0

	// 模拟并发更新，启动50个线程，对计数器进行+1操作
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 每隔线程无限循环让计数器增加（操作内存地址）
				atomic.AddUint64(&ops, 1)

				//允许其他go routine的执行
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	// 为了在计数器还在被其他go routine更新时安全的使用它，通过atomic.LoadUint64()将当前值的拷贝提取到opsFinal中
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)

}
