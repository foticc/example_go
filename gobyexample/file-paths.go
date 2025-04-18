package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("d:\\", "a", "b", "c.jpg")
	fmt.Println(p)

	fmt.Println(filepath.Join("/sfasf", "sdfa", "../sdf"))

	fmt.Println(filepath.Base(p))
	fmt.Println(filepath.Dir(p))
	fmt.Println(filepath.Split(p))

	ext := filepath.Ext(p)
	fmt.Println(ext)

	fmt.Println(filepath.IsAbs("/a/b/v.txt"))
	fmt.Println(filepath.IsAbs("c:\\a\\b\\v.txt"))

	fmt.Println(strings.TrimSuffix(filepath.Base(p), ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/d", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
