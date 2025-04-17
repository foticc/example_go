package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := os.WriteFile("example.txt", []byte("Hello,World!"), 0644)
	check(err)

	f, err := os.Create("example2.txt")
	check(err)
	defer f.Close()

	l1, err := f.Write([]byte{12, 12, 14})
	check(err)
	fmt.Println(l1)

	l2, err := f.WriteString("Hello,World!")
	check(err)
	fmt.Println(l2)
	f.Sync()

	w := bufio.NewWriter(f)
	l3, err := w.WriteString("WriteString")
	check(err)
	fmt.Println(l3)
	w.Flush()
}
