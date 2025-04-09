package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	fmt.Println(len(s))

	s[0] = "0"
	s[1] = "1"
	s[2] = "2"
	fmt.Println("s=", s)
	s = append(s, "3")
	fmt.Println("s=", s)

	c := make([]string, 4)
	copy(c, s)
	fmt.Println("c=", c)

	fmt.Println("s[:2]=", s[:2])
	fmt.Println("s[1:]=", s[1:])
	fmt.Println("s[1:2]=", s[0:2])

}
