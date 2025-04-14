package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Println("work start", id)
	time.Sleep(time.Second)
	fmt.Println("work done", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
	fmt.Println("all done")
}
