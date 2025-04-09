package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("arr = ", arr)
	arr[0] = 10
	arr[4] = 20
	fmt.Println("set 0,4 = ", arr)
}
