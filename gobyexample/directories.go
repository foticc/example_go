package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := os.Mkdir("mydir", 0755)
	check(err)
	// rm -rf
	defer os.RemoveAll("mydir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("mydir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	check(err)

	for _, f := range c {
		fmt.Println(f.Name(), f.IsDir())
	}
	// Chdir 可以修改当前工作目录，类似于 cd。
	err = os.Chdir("subdir/parent/child")
	check(err)

	err = os.Chdir("../../..")
	check(err)

	filepath.Walk("subdir", func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(" ", p, info.IsDir())
		return nil
	})

}
