package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"elderberry", "apple", "banana", "cherry", "date", "123"}

	sort.Strings(strs)

	fmt.Println(strs)

	ints := []int{5, 2, 8, 3, 1, 9}

	fmt.Println(sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Println(sort.IntsAreSorted(ints))
	fmt.Println(ints)
}
