package example10

import (
	"fmt"
	"testing"
)

// 通道遍历
func TestRangeOverChannel(t *testing.T) {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// range迭代queue中得到的每个值。因为close，这个迭代会在2个值之后结束，如果没有close，将在这个循环继续阻塞，等待接收
	for elem := range queue {
		fmt.Println(elem)
	}
}
