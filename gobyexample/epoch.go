package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	secs := t.Unix()
	nanos := t.UnixNano()
	fmt.Println(secs)
	fmt.Println(nanos)

	millis := nanos / 1000000
	fmt.Println(millis)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))

}
