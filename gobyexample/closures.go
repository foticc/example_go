package main

import "fmt"

func autoIncrement() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	gen := autoIncrement()

	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())

	gen1 := autoIncrement()
	fmt.Println(gen1())
	fmt.Println(gen1())
}
