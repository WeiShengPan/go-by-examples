package singleton

import "sync"

// 单例实现，使用sync.Once
type Singleton struct {
}

var instance *Singleton
var tmOnce sync.Once

func GetTaskConsumer() *Singleton {
	// sync.Once会保证Do中的方法只被执行一次
	tmOnce.Do(func() {
		instance = &Singleton{}
		instance.init()
	})
	return instance
}

func (instance *Singleton) init() {
	// do init things...
}
