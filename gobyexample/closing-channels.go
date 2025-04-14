package main

import "fmt"

func main() {
	jobs := make(chan string, 6)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job:", j)
			} else {
				fmt.Println("no more jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 6; i++ {
		fmt.Println("send job", i)
		jobs <- fmt.Sprintf("job %d", i)
	}

	close(jobs)
	fmt.Println("sent all jobs")
	<-done

}
