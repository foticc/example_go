package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)
	for i, a := range args {
		fmt.Printf("arg %d: %s\n", i, a)
	}
}
