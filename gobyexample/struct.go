package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func main() {
	p := Person{name: "John", age: 30}
	fmt.Println(p.name, p.age)

	fmt.Println(newPerson("Alice", 18))
	fmt.Println(&Person{name: "John", age: 30})

	a := Person{name: "Alice", age: 18}
	fmt.Println(a.name, a.age)
}
