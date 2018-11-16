package example11

import (
	"fmt"
	"testing"
	"time"
)

// ticker 打点器，在固定的时间间隔重复执行，将会定时执行，知道被停止
func TestTickers(t *testing.T) {

	ticker := time.NewTicker(time.Second)

	// ticker.C 的通道应用于发送数据
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// ticker一旦停止，将不能再从它的通道中接收到值
	time.Sleep(time.Second * 5)
	ticker.Stop()
	fmt.Println("Ticker stopped")

	time.Sleep(time.Second * 3)
}
