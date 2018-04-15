package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	searchIn := "John: 3289.34 Willam: 43.30 stve: 93044.10"
	pat := "[0-9]+.[0-9]+"

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found")
	}

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

}
