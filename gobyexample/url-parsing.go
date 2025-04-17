package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?allow=true#f"

	r, e := url.Parse(s)
	if e != nil {
		panic(e)
	}
	fmt.Println(r.Scheme)
	fmt.Println(r.Host)
	fmt.Println(r.Path)
	fmt.Println(r.User.Username())
	fmt.Println(r.User.Password())
	fmt.Println(r.User.String())
	fmt.Println(r.Port())

	fmt.Println(r.Fragment)
	fmt.Println(r.RawQuery)

	m, _ := url.ParseQuery(r.RawQuery)
	fmt.Println(m)
}
