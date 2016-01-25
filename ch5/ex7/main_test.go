package main

import (
	"testing"
)

func TestOutline(t *testing.T) {
	if err := outline("https://golang.org"); err != nil {
		t.Fatalf("outline error: %s")
	}
}
