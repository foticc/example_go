package main
import "fmt"

const uri = "http://www.example.com"

type Person struct {
	name string
	age int
}

var array [5]int = [5]int{1,3,4,5,6}

func main() {
	const (
		n1 = iota
		n2
		n3
	)
	s1:= "❤"
	var a,b,c int = 1,2,3
	fmt.Println(n1, n2, n3)
	fmt.Println(s1)
	fmt.Println(a,b,c)
	fmt.Println(uri)
	fmt.Println(array)
	Person1 := Person{"Alice", 25}
	fmt.Println(Person1)
}