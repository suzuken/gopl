package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"gopl.io/ch5/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(rawurl string) func(item string) []string {
	fmt.Println(rawurl)
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Print(err)
	}
	tempdir := os.TempDir()
	return func(item string) []string {
		fmt.Println(item)
		itemUrl, err := url.Parse(item)
		if err != nil {
			log.Print(err)
		}
		// if same host, save contents on disk.
		if itemUrl.Host == u.Host {
			file, err := ioutil.TempFile(tempdir, "crawler-")
			defer file.Close()
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("write %s to %s\n", item, file.Name())
			resp, err := http.Get(item)
			defer resp.Body.Close()
			if err != nil {
				log.Print(err)
			}
			if _, err := io.Copy(file, resp.Body); err != nil {
				log.Print(err)
			}
		}
		list, err := links.Extract(item)
		if err != nil {
			log.Print(err)
		}
		return list
	}

}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl(os.Args[1]), os.Args[1:])
}
