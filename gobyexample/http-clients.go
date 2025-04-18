package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Status:", res.Status)

	scanner := bufio.NewScanner(res.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
