package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Float64() * 100)

	rn := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(rn)

	fmt.Println(r1.Intn(100))

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
