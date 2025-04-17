package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "base64 encoding!&@(#@*&$!@**~)"

	bse := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(bse)

	bsd, _ := base64.StdEncoding.DecodeString(bse)
	fmt.Println(string(bsd))

	bue := base64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println(bue)

	bud, _ := base64.URLEncoding.DecodeString(bue)
	fmt.Println(string(bud))
}
