package main

import (
	"fmt"
	"sort"
	"testing"
)

// 使用函数自定义排序
func TestSortingByFunctions(t *testing.T) {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

// 为了在 Go 中使用自定义函数进行排序，我们需要一个对应的
// 类型。这里我们创建一个为内置 `[]string` 类型的别名的
// `ByLength` 类型
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

//按字符串长短排序
func (s ByLength) Less(i int, j int) bool {
	return len(s[i]) < len(s[j])
}
