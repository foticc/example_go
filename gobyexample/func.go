package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func add3(a, b int) (int, int) {
	return a + b, a - b
}

func sum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("Hello, world!")
	c := add(1, 2)
	fmt.Println(c)

	s, s1 := add3(1, 2)

	fmt.Println(s, s1)

	result := sum(1, 2, 31, 2, 41, 4, 1, 41, 2, 1, 5, 1, 51, 5, 1, 51, 51, 5)
	fmt.Println(result)

	r := fact(5)
	fmt.Println(r)
}
