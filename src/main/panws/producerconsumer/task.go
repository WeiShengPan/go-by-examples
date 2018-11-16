package producerconsumer

import (
	"fmt"
)

type Task struct {
	Id   int
	Name string
}

func (t *Task) Print() {
	fmt.Println("id:", t.Id, ", name:", t.Name)
}
