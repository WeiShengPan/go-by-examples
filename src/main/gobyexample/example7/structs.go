package example7

import "fmt"

// 结构体
func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{age:30,name: "Alice"})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{"Ann", 40})

	s:=person{"Sean", 50}
	fmt.Println(s.name)

	sp:=&s
	fmt.Println(sp.age)

	sp.age= 51
	fmt.Println(s)
	
}

type person struct {
	name string
	age int
}
