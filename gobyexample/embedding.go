package main

import "fmt"

type base struct {
	num int
}

func (b base) desc() string {
	return fmt.Sprintf("base num: %d", b.num)
}

type container struct {
	base
	name string
}

func main() {
	c := container{
		base: base{
			num: 10,
		},
		name: "哈哈",
	}
	fmt.Println("container:", c)
	fmt.Println("num:", c.base.num)
	fmt.Println("num:", c.num)
	fmt.Println("name:", c.name)

	type describer interface {
		desc() string
	}
	var d describer = c
	fmt.Println("describer:", d.desc())
	fmt.Println("describer:", d)
}
