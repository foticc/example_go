package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	b, e := os.ReadFile("test.json")
	check(e)
	fmt.Println(string(b))

	fo, e := os.Open("test.json")
	defer fo.Close()
	check(e)
	bm := make([]byte, 15)
	l, e := fo.Read(bm)
	check(e)
	fmt.Println(l, string(bm[:l]))

	o2, e := fo.Seek(15, 0)
	check(e)
	b2 := make([]byte, 30)
	n2, err := fo.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := fo.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(fo, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = fo.Seek(16, 0)
	check(err)

	r4 := bufio.NewReader(fo)
	b4, err := r4.Peek(10)
	check(err)
	fmt.Println(string(b4))
}
