package main

import "fmt"

func pings(pings chan<- string, msg string) {
	pings <- msg
}

func pongs(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	i := make(chan string, 1)
	o := make(chan string, 1)

	pings(i, "123")
	pongs(i, o)

	fmt.Println(<-o)
}
