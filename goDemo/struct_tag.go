package main

import (
	"fmt"
	"reflect"
)

type TagType struct { //tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := TagType{true, "barrak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, i int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(i)
	fmt.Printf("%v \n", ixField.Tag)
}
