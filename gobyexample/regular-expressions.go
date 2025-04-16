package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	r, _ := regexp.MatchString(`^\w+\d*[mol]$`, "k123314124m")

	fmt.Println(r)

	p, _ := regexp.Compile(`\w+\d*[mol]`)

	fmt.Println(p.MatchString("hhah"))
	fmt.Println(p.FindString("k1mfassdf"))
	fmt.Println(p.FindStringIndex("k1mfassdf"))

	fmt.Println(p.FindStringSubmatch("k1mfassq3l"))
	fmt.Println(p.FindStringSubmatchIndex("k1mfassq3l"))

	fmt.Println(p.FindAllStringSubmatchIndex("k1mfassq3l", -1))

	fmt.Println(p.FindAllString("k1mfassq3l", 1))

	fmt.Println(p.Match([]byte("df8")))

	rm := regexp.MustCompile(`[a-h]`)
	fmt.Println(rm)
	fmt.Println(rm.FindAllString("z", -1))

	fmt.Println(rm.ReplaceAllString("hello world", "X"))

	out := rm.ReplaceAllStringFunc("hello world", func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Println(out)
}
