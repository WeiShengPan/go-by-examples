package producerconsumer

import (
	"sync"
)

type TaskCache struct {
	waitingQueue chan *Task
}

var taskCache *TaskCache
var rOnce sync.Once

// singleton
func GetTaskCache() *TaskCache {
	rOnce.Do(func() {
		taskCache = &TaskCache{}
		taskCache.waitingQueue = make(chan *Task, 5)
	})
	return taskCache
}

func (cache *TaskCache) Push(t *Task) {
	if t == nil {
		return
	}
	cache.waitingQueue <- t
}

func (cache *TaskCache) Pop() *Task {
	t := <-cache.waitingQueue
	return t
}
