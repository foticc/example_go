package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	var b int32 = 123
	// for i := 1; i < len(os.Args); i++ {
	// 	 s+=os.Args[i]+"\n"
		
	// }
	s = strings.Join(os.Args[0:], "\n")
	fmt.Println(s)
	fmt.Println(b)
}