package main

import (
	"flag"
	"fmt"
)

// go build command-line-flags.go
// command-line-flags -word=text -numb=1 -name=hha -fork=false hello hi
func main() {
	prt := flag.String("word", "foo", "str")
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "name", "default", "a string var")
	// Must be called after all flags are defined and before flags are accessed by the program.
	flag.Parse()
	fmt.Println(*prt)
	fmt.Println(*numbPtr)
	fmt.Println(*forkPtr)
	fmt.Println(svar)
	fmt.Println(flag.Args())

}
