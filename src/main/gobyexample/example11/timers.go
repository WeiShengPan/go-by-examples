package example11

import (
	"fmt"
	"time"
)

// timer 定时器
func main() {

	// 定时器将等待2s
	timer1:=time.NewTimer(time.Second* 2)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 可以在定时器失效之前取消这个定时器。这里在go routine 中等待阻塞，主线程中直接stop，所以timer2不会失效，直接停止
	timer2:=time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2:=timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
