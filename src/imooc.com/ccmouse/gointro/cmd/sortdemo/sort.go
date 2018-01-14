package main

import (
	"fmt"
	"sort"
)

func main() {
	// create a slice of int
	a := []int{3, 6, 7, 6, 1, 10, 8}

	sort.Ints(a)

	for _, v := range a {
		fmt.Println(v)
	}

}
