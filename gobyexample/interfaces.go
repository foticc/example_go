package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return (r.width + r.height) * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("geometry: ", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}

func main() {
	r := rectangle{width: 10, height: 5}
	c := circle{radius: 3}

	measure(r)
	measure(c)
}
