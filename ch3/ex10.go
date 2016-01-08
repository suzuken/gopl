package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	for i, v := range s {
		if (n-i)%3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(v)
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("10000000"))
	fmt.Println(comma("1000000000"))
}
