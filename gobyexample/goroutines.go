package main

import (
	"fmt"
	"time"
)

func f(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str, i)
	}
}

func main() {
	go f("hello")

	go func() {
		fmt.Println("123")
	}()
	time.Sleep(time.Second)
	fmt.Println("world!")
}
