package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "usage: ./ex5.go http://example.com\n")
		os.Exit(1)
	}
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("words: %d, images: %d\n", words, images)

}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			words++
		}
	} else if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cw, ci := countWordsAndImages(c)
		words = words + cw
		images = images + ci
	}
	return
}
