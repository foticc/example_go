package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Path    string   `xml:"path"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant: %s, %s", p.Name, p.Path)
}

func main() {
	p := &Plant{
		Id:   123,
		Name: "John",
		Path: "C:\\Users\\John\\Documents",
	}
	out, _ := xml.MarshalIndent(p, " ", "  ")
	fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))

	var p1 Plant
	if err := xml.Unmarshal(out, &p1); err != nil {
		panic(err)
	}
	fmt.Println(p1)

	type Nesting struct {
		XMLName xml.Name `xml:"Nesting"`
		Plants  []Plant  `xml:"parent>child>plant"`
	}
	n := Nesting{
		Plants: []Plant{
			*p,
		},
	}
	out, _ = xml.MarshalIndent(n, " ", "  ")
	fmt.Println(string(out))
}
