package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Java path:", os.Getenv("JAVA_HOME"))
	fmt.Println("Cargo path:", os.Getenv("CARGO_HOME"))
	fmt.Println()

	for _, env := range os.Environ() {
		s := strings.SplitN(env, "=", 2)
		fmt.Println(s[0])
	}
}
