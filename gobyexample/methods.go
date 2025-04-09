package main

import "fmt"

type Rectange struct {
	width  int
	height int
}

func (r *Rectange) area() int {
	return r.width * r.height
}

func area(r Rectange) int {
	return r.width * r.height
}

func main() {

	r := Rectange{width: 10, height: 20}
	fmt.Println("area=", area(r))
	fmt.Println("area=", r.area())

	rp := &r
	fmt.Println("area=", rp.area())

}
