package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("contains: ", s.Contains("hello world", "l"))
	p("Count:    ", s.Count("hello world", "l"))
	p("HasPrefix:", s.HasPrefix("hello world", "l"))
	p("HasSuffix:", s.HasSuffix("hello world", "d"))
	p("Index     ", s.Index("hello world", "l"))
	p("Join      ", s.Join([]string{"1", "2", "3"}, "-"))
	p("Repeat    ", s.Repeat("hi", 3))
	p("Replace   ", s.Replace("hello world", "l", "L", -1))
	p("Replace   ", s.Replace("hello world", "l", "L", 1))
	p("Split     ", s.Split("a,b,c", ","))
	p("ToLower   ", s.ToLower("HELLO WORLD"))
	p("ToUpper   ", s.ToUpper("hello world"))
	p("TrimSpace ", s.TrimSpace("  hello world  "))
	p()
	p("Len       ", len("hello world"))
	p("char      ", "哈哈哈"[0])
}
