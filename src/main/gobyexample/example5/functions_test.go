package example5

import (
	"fmt"
	"testing"
)

// 函数
func TestFunctions(t *testing.T) {

	res1 := plus(1, 2)
	fmt.Println(res1)

	res2, res3, res4 := values(1, 2, 3)
	fmt.Printf("%d %d %d\n", res2, res3, res4)

	res5 := sum(1, 2, 3, 4)
	fmt.Println(res5)
}

// 单返回值
func plus(a int, b int) int {
	return a + b
}

// 多返回值
func values(a int, b int, c int) (int, int, int) {
	return a + 1, b + 1, c + 1
}

// 变参函数
func sum(nums ...int) int {

	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
