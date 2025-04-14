package main

import (
	"sync"
	"sync/atomic"
)

func main() {

	var count uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for w := 0; w < 1000; w++ {
				atomic.AddUint64(&count, 1)
			}
			defer wg.Done()
		}()
	}

	wg.Wait()
	println("Final count is:", count)

}
