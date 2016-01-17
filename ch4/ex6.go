package main

import (
	"fmt"
	"unicode"
)

func convert(ss string) string {
	var runes []rune
	for _, s := range ss {
		if unicode.IsSpace(s) {
			runes = append(runes, ' ')
		} else {
			runes = append(runes, s)
		}
	}
	return string(runes)
}

func main() {
	x := "aa\nkuke\rhoge"
	fmt.Println(convert(x))
}
