package example7

import (
	"fmt"
	"math"
)

// 接口是方法特征的命名集合
func main() {

	s := square{3, 4}
	measure(s)

	c := circle{5}
	measure(c)
}

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
