package main

import (
	"fmt"
)

func exists(needle string, ss []string) bool {
	for _, s := range ss {
		if s == needle {
			return true
		}
	}
	return false
}

func dedup(ss []string) []string {
	out := ss[:1]
	for _, s1 := range ss {
		if !exists(s1, out) {
			out = append(out, s1)
		}
	}
	return out
}

func main() {
	x := []string{"aa", "aa", "bb", "cc", "bb"}
	fmt.Println(dedup(x))
}
