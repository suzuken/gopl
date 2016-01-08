package main

import (
	"bytes"
	"fmt"
)

// anagram returns anagrem of string
func anagram(s string) string {
	var buf bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func isAnagram(a, b string) bool {
	return anagram(a) == b
}

func main() {
	fmt.Println(anagram("abcde"))
	fmt.Println(isAnagram("abcde", "edcba"))
	fmt.Println(isAnagram("abcde", "kuke"))
}
