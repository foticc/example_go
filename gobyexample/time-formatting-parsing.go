package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))
	p(t.Format(time.DateTime))

	t1, _ := time.Parse(time.RFC3339, "2025-04-17T15:22:22+08:00")
	p(t1)

	p(t1.Format(time.Kitchen))

	form := "3 04 PM"

	t2, e := time.Parse(form, "8 41 PM")
	p(t2)
	p(e)
	ansic := "Mon Jan _2 15:04:05 2006"
	_, err := time.Parse(ansic, "8:41PM")
	p(err)
}
