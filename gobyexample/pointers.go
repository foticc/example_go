package main

import "fmt"

func zeroval(x int) {
	fmt.Println("zeroval called", x)
	x = 123
}

func zeroptr(x *int) {
	fmt.Println("zeroptr called", *x)
	*x = 0
}

func main() {

	a := 123
	zeroptr(&a)
	zeroval(a)
	zeroval(a)
}
