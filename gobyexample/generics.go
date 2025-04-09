package main

import "fmt"

func add[T int | float64](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println("Hello, world!")
	a := add(1, 2)
	b := add(1.5, 2.5)
	fmt.Println(a, b)
}
