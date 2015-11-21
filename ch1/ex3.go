package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Now().Sub(t1))

	t2 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Now().Sub(t2))
}
