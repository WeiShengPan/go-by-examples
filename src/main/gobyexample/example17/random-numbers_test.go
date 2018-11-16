package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// 随机数
func TestRandomNumbers(t *testing.T) {

	var p = fmt.Println

	p(rand.Intn(100))

	// 返回一个浮点数 0.0<=f<=1.0
	p(rand.Float64())

	// 返回 5.0<=f<=10.0
	p((rand.Float64() * 5) + 5)

	// 给定种子
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	p(r1.Intn(100))

	// 相同的种子，会产生相同的随机数序列
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	p(r2.Intn(100))

}
