package main

import (
	"fmt"
	"time"
)

func main() {
	// base limiter***********
	requests := make(chan int, 6)

	for i := 1; i <= 6; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(1 * time.Second)

	for req := range requests {
		<-limiter
		fmt.Println("requests", req, time.Now())
	}
	// base limiter***********

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for r := range time.Tick(1 * time.Second) {
			burstyLimiter <- r
		}
	}()

	burstyRequests := make(chan int, 6)
	for i := 1; i <= 6; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
