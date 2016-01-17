package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
	// [2 3 4 5 0 1]
}

func rotate(x []int, i int) {
	tmp := make([]int, len(x))
	copy(tmp, x)
	copy(x, x[i:])
	copy(x[len(x)-i:], tmp[0:i])
}
