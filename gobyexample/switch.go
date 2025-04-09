package main

import "fmt"

func main() {
	switch num := 10; {
	case num < 0:
		fmt.Println("Negative number")
	case num == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Positive number")
	}
	n := 10
	switch {
	case n > 0:
		fmt.Println("Positive number")
	case n == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Negative number")
	}
}
