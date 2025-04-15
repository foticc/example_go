package main

import (
	"fmt"
	"sort"
)

type newSortLen []string

func (c newSortLen) Len() int {
	return len(c)
}

func (c newSortLen) Swap(i1, i2 int) {
	c[i1], c[i2] = c[i2], c[i1]
}

func (c newSortLen) Less(i1, i2 int) bool {
	return len(c[i1]) < len(c[i2])
}

func main() {
	s := []string{"apple", "banana", "chy", "da", "elderberry"}
	sort.Sort(newSortLen(s))
	fmt.Println(s)
}
