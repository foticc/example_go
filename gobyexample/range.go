package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for index, value := range nums {
		fmt.Println("inex", index, "value", value)
		sum += value
		fmt.Println("current sum", sum)
	}
	fmt.Println("final sum", sum)

	kvs := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(kvs)
	for key, value := range kvs {
		fmt.Println(key, value)
	}

	for key := range kvs {
		fmt.Println(key)
	}

	for i, s := range "h😘💕😒😒😒😁😁😊😎❤️" {
		fmt.Println(i, s)
	}
}
