package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 18})
	fmt.Println(person{"jimmy", 30})
	fmt.Println(person{name: "jim"})
	fmt.Println(person{age: 23})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println("sp.age:", sp.age)
	sp.age = 51
	fmt.Println("sp.age:", sp.age)
	fmt.Println("s.age:", s.age)

	ss := s
	ss.age = 60
	fmt.Println("ss.age", ss.age)
	fmt.Println("s.age", s.age)
}
