package main

import (
	"./person"
	"fmt"
)

func main() {
	p := new(person.Person)
	p.SetFirstName("erci")
	fmt.Println(p.FirstName())
}
