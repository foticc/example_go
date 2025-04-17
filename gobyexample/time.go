package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main() {
	t := time.Now()
	p("Hello, world!", t)

	td := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	p(td)
	p(td.Year())
	p(td.Month())
	p(td.Day())
	p(td.Hour())
	p(td.Minute())
	p(td.Second())
	p(td.Nanosecond())

	p(td.Weekday())

	p(td.Before(t))
	p(td.After(t))
	p(td.Equal(t))

	diff := td.Sub(t)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(td.Add(diff))
	p(td.Add(-diff))

	p(td.Add(2 * time.Hour))

}
