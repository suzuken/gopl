package main

import (
	"fmt"
)

// it works and not allocaate new memory, but only support 3 byte unicode character..
func reverse(b []byte) {
	for i, j := 0, len(b)-3; i < j; i, j = i+3, j-3 {
		b[i], b[j] = b[j], b[i]
		b[i+1], b[j+1] = b[j+1], b[i+1]
		b[i+2], b[j+2] = b[j+2], b[i+2]
	}
}

func main() {
	s := []byte("こんにちは")
	reverse(s)
	fmt.Println(string(s))
}
