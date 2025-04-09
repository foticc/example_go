package main

import "fmt"

func main() {
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0

	for {
		fmt.Println(j)
		if j > 5 {
			break
		}
		j++
	}
}
