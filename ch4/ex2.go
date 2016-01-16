package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var bit = flag.Int("bit", 256, "bit of SHA. Default: 256.")

func main() {
	flag.Parse()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	switch *bit {
	case 256:
		fmt.Fprintf(os.Stdout, "%x", sha256.Sum256(b))
	case 384:
		fmt.Fprintf(os.Stdout, "%x", sha512.Sum384(b))
	case 512:
		fmt.Fprintf(os.Stdout, "%x", sha512.Sum512(b))
	default:
		fmt.Fprint(os.Stderr, "unsupported bit")
	}
}
