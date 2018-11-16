package producerconsumer

import (
	"testing"
	"time"
)

// 模拟生产者消费者问题
func TestProducerConsumer(t *testing.T) {

	// 消费者先等待
	GetTaskConsumer()

	time.Sleep(time.Second * 5)

	// 生产者
	GetTaskProducer().Produce()
}
