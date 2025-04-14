package main

import "fmt"

func main() {

	msg_chan := make(chan string, 2)

	go func() {
		msg_chan <- "message from goroutine"
		msg_chan <- "message from goroutine2"
	}()

	for i := 0; i < 2; i++ {
		msg := <-msg_chan

		fmt.Println(msg)
	}

}
