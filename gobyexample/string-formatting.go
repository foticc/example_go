package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y float32
}

var pf = fmt.Printf

func main() {
	p := point{1.123, 2.123}
	pf("struct: %v\n", p)
	pf("struct: %+v\n", p)
	pf("struct: %#v\n", p)

	pf("type:   %T\n", p)
	pf("bool:   %t\n", true)
	pf("int:    %d\n", 100)

	pf("bin:    %b\n", 100)
	pf("char:   %c\n", 66)
	pf("hex:    %x\n", 314159123)

	pf("floatf:  %f\n", 3.141592653589793)
	pf("floate:  %e\n", 3.141592653589793)
	pf("floatE:  %E\n", 3.141592653589793)

	pf("string:  %s\n", "string")
	pf("string:  %q\n", "string")
	pf("string:  %x\n", "string")

	pf("pointer: %p\n", &p)

	pf("width1: |%6d|%3d|\n", 100, 21)
	pf("width2: |%6.2f|%6.2f|\n", 1.2, 3.4562)
	pf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	pf("width4: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("Sprintf: %s\n", "foo")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}
