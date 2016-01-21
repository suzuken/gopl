package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex.5-2: %v\n", err)
		os.Exit(1)
	}
	m := make(map[string]int)
	fmt.Printf("%#v\n", elementCount(m, doc))
}

func elementCount(em map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		em[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elementCount(em, c)
	}
	return em
}
