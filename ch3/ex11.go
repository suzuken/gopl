package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		return comma(s[:dot]) + s[dot:]
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(comma("10000000"))
	fmt.Println(comma("10000000.0000000"))
}
