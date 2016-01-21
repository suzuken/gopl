package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "5-4: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "link":
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Printf("%s\n", a.Val)
				}
			}
		case "script", "img":
			for _, a := range n.Attr {
				if a.Key == "src" {
					fmt.Printf("%s\n", a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
