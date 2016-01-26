package main

import (
	"fmt"
)

func nonzero() {
	panic(1)
}

func main() {
	defer func() {
		v := recover()
		fmt.Println(v)
	}()
	nonzero()
}
