package producerconsumer

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 生产者
type TaskProducer struct {
}

var producer *TaskProducer
var tcOnce sync.Once

func GetTaskProducer() *TaskProducer {
	tcOnce.Do(func() {
		producer = &TaskProducer{}
	})
	return producer
}

// 生产
func (tc *TaskProducer) Produce() {
	i := 0
	for {
		fmt.Println("Produce task :", i)
		task := new(Task)
		task.Id = i
		task.Name = "Task " + strconv.Itoa(i)
		GetTaskCache().Push(task)
		i++
		time.Sleep(time.Second * 1)
	}
}
