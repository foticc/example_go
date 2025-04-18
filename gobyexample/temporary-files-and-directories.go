package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp("", "temp")
	check(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte("Hello, world!"))
	check(err)

	dname, err := os.MkdirTemp("", "tmp")

	check(err)
	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file.txt")
	err = os.WriteFile(fname, []byte("Hello, world!"), 0666)
	check(err)
}
