package main

import (
	"errors"
	"fmt"
)

func add(sum int) (int, error) {
	if sum == 1 {
		return -1, errors.New("errors sum")
	}
	return sum, nil
}

func main() {
	sum, err := add(0)
	if err != nil {
		fmt.Printf("error:\n")
	} else {
		fmt.Printf("sum: %v\n", sum)
	}
}
