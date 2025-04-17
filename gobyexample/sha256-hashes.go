package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "som string to hash"

	sn := sha256.New()
	sn.Write([]byte(s))
	hashed := sn.Sum(nil)

	fmt.Printf("%x\n", hashed)
}
