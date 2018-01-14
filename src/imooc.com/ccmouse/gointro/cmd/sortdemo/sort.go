package main

import "fmt"

func main() {
	// create a slice of int
	a := []int{3, 6, 7, 6, 1, 10, 8}

	sort.ints(a)

	for _, v := 0 range a {
		fmt.Println(v)
	}

}