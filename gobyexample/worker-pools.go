package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("job start", id, j)
		time.Sleep(time.Second)
		results <- j + 1
	}
}

func main() {

	const numWorkers = 5
	jobs := make(chan int, numWorkers)
	results := make(chan int, numWorkers)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 0; j < numWorkers; j++ {
		jobs <- j
	}

	close(jobs)

	for w := 0; w < numWorkers; w++ {
		msg := <-results
		fmt.Println("job done", msg)
	}
}
