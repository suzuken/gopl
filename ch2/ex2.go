package main

import (
	"fmt"
	"github.com/suzuken/gopl/ch2/weightconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cw: %v\n", err)
			os.Exit(1)
		}
		k := weightconv.Kilogram(t)
		p := weightconv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n", k, weightconv.KToP(k), p, weightconv.PToK(p))
	}
}
