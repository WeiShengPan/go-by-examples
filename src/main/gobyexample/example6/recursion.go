package example6

import "fmt"

// 递归
func main() {
	fmt.Println(fact(7))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
