package main

import "fmt"

type newType[T int, T2 float64] struct {
	value  T
	value2 T2
}

func add[T int | float64](a T, b T) T {
	return a + b
}

func cout[T fmt.Stringer](str T) {
	fmt.Println(str.String())
}

func main() {
	fmt.Println("Hello, world!")
	a := add(1, 2)
	b := add(1.5, 2.5)
	fmt.Println(a, b)

	newType1 := newType[int, float64]{value: 1, value2: 2.5}
	fmt.Println("newType1:", newType1)

}
