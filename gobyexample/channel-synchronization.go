package main

import (
	"fmt"
	"time"
)

func wait(done chan bool) {
	fmt.Println("start")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {

	done := make(chan bool)
	go wait(done)
	<-done
}
