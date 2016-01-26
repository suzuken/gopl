package main

import (
	"fmt"
)

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	min := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	fmt.Println(max(1, 3, 10, 2, 100))
	fmt.Println(min(1, 3, 10, 2, 100))

	fmt.Println(max())
	// 0
}
