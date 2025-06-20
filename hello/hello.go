package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {

	students := [2]Student{Student{Name: "Alice", Age: 20}, Student{Name: "Amilia", Age: 18}}
	fmt.Printf("%+v", students)
}
