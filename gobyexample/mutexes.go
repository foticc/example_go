package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {

	container := Container{
		counters: map[string]int{"a": 0, "b": 0, "c": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, count int) {
		for i := 0; i < count; i++ {
			container.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 100000)
	go doIncrement("b", 100000)
	go doIncrement("c", 100000)

	wg.Wait()

	fmt.Println(container.counters)

}
