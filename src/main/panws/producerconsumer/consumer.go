package producerconsumer

import (
	"fmt"
	"sync"
)

// 消费者
type TaskConsumer struct {
}

var taskConsumer *TaskConsumer
var tmOnce sync.Once

func GetTaskConsumer() *TaskConsumer {
	tmOnce.Do(func() {
		taskConsumer = &TaskConsumer{}
		taskConsumer.init()
	})
	return taskConsumer
}

func (tm *TaskConsumer) init() {
	// 这里需要开另外线程去消费，否则会造成死锁
	go tm.Consume()
}

// 消费
func (tm *TaskConsumer) Consume() {
	for {
		fmt.Println("Begin consume...")
		t := GetTaskCache().Pop()
		fmt.Println("Consume task:", t)
	}
}
