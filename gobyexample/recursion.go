package main

import "fmt"

func main() {
	var f func(n int) int
	f = func(n int) int {
		if n < 2 {
			return n
		}
		return f(n-1) + f(n-2)
	}
	sum := f(10)
	fmt.Println(sum)
}
