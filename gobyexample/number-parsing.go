package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)

	h, _ := strconv.ParseInt("0x12f", 0, 64)
	fmt.Println(h)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	a, _ := strconv.Atoi("135")
	fmt.Println(a)

	_, e := strconv.Atoi("error")
	fmt.Println(e)
}
