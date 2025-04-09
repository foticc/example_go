package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "你是谁"
	fmt.Println(s)

	fmt.Println("len=", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Println("c", s[i])
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, c := range s {
		fmt.Printf("Rune at index %d: %c\n", idx, c)
	}
}
