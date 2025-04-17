package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.ToUpper(scanner.Text())
		fmt.Println(">>", text)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}
