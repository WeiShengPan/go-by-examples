package main

import (
	"container/list"
	"fmt"
	"testing"
)

/**
 * 继承
 * 实现方式：一个struct A嵌套另一个匿名struct B，那么这个A可以直接访问匿名B的方法，实现继承；如果一个struct A嵌套了多个
 * 匿名struct B、struct C，那么这个A可以直接访问多个匿名struct的方法(B或C的所有方法)，实现多重继承
 *
 * --- 与组合的区别：struct A中嵌入的是另一个有名的struct B，这是组合，A中不能访问B中的方法 ---
 *
 * 与interface搭配使用，只要父类实现了interface的所有方法，子类也会自动拥有父类所实现的interface所有方法，父类和子类都可
 * 以转化为对应的interface
 */
func TestInheritance(t *testing.T) {
	taskList := list.New()
	taskList.PushBack(&Task{TaskId: 1, Result: 1})
	taskList.PushBack(&TaskA{Task: Task{TaskId: 2, Result: 2}, aName: "A"})
	taskList.PushBack(&TaskB{Task: Task{TaskId: 3, Result: 3}, bName: "B"})
	taskList.PushBack(&TaskC{cName: "c"})

	for ele := taskList.Front(); ele != nil; ele = ele.Next() {
		//这里不管是父类Task，还是子类TaskA TaskB，都可以转为TaskInterface，因为TaskA TaskB继承了Task
		//if t, ok := (ele.Value).(TaskInterface); ok {
		//	fmt.Println(t.GetTaskId(), t.GetResult())
		//} else {
		//	fmt.Println("Invalid element in task list:", ele.Value)
		//}
		t := (ele.Value).(*TaskA)
		if t != nil {
			fmt.Println(t)
		}
	}
}

type TaskInterface interface {
	GetTaskId() int
	GetResult() int
}

type Task struct {
	TaskId int
	Result int
}

func (task *Task) GetTaskId() int {
	return task.TaskId
}

func (task *Task) GetResult() int {
	return task.Result
}

type TaskA struct {
	Task
	aName string
}

type TaskB struct {
	Task
	bName string
}

type TaskC struct {
	cName string
}
