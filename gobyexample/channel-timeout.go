package main

import (
	"time"
)

func main() {
	ch1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "hello"
	}()

	select {
	case msg := <-ch1:
		println(msg)
	case <-time.After(1 * time.Second):
		println("timeout")
	}

	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "world"
	}()

	select {
	case msg := <-ch2:
		println(msg)
	case <-time.After(3 * time.Second):
		println("timeout")
	}

}
