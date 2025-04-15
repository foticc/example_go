package main

import "fmt"

func mayPanic() {
	panic("may have some panic")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	mayPanic()
	fmt.Println("hello")

}
