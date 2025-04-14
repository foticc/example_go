package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		fmt.Println("Waiting for Timer 2 to expire")
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop := timer2.Stop()

	if stop {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
