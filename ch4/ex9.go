package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq() {
	words := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	for k, c := range words {
		fmt.Printf("word: %s\tcount: %d\n", k, c)
	}
}

func main() {
	wordfreq()
}
