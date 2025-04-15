package main

import (
	"fmt"
	"os"
)

func createFile(path string) *os.File {
	// file, err := os.Create(path)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return file
}

func writeToFile(file *os.File, data string) {
	fmt.Fprintln(file, data)
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("close error")
		os.Exit(1)
	}
}

func main() {
	f := createFile("test.txt")
	defer closeFile(f)
	writeToFile(f, "Hello, world!fasfsadf")
}
