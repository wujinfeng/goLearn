package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.1415956
	mapAssigned["two"] = 3

	fmt.Printf("map litrral at one is : %d\n", mapLit["one"])
	fmt.Printf("map created at key2 is : %f\n", mapCreated["key2"])
	fmt.Printf("map assigned at two is : %d\n", mapLit["two"])
	fmt.Printf("map litrral at ten is : %d\n", mapLit["ten"])
}
