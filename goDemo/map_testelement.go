package main

import "fmt"

func main() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["washington"] = 32
	value, isPresent = map1["Beijing"]
	if isPresent {
		fmt.Printf("the value of Beijing in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("is Paris in map1 ?: %t \n", isPresent)
	fmt.Printf("value is : %d\n ", value)

	delete(map1, "washington")
	value, isPresent = map1["washington"]
	if isPresent {
		fmt.Printf("the value of washington in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain washington")
	}

}
