package main

import (
	"fmt"
	"sort"
	"testing"
)

// 排序
// 内置数据类型的排序
func TestSorting(t *testing.T) {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	// 检查一个序列是不是已经排好序
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
