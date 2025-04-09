package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["hello"] = "world"
	m["name"] = "who"
	fmt.Println(m)
	fmt.Println("hello:", m["hello"])

	// 值，是否存在key
	s, s1 := m["name"]
	fmt.Println(s, s1)
}
