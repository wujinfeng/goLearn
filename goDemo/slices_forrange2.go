package main

import "fmt"

func main() {
	seasons := []string{"spring", "summer", "autumn", "winter"}
	for ix, season := range seasons {
		fmt.Printf("season %d is: %s\n", ix, season)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
}
