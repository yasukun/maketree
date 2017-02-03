package main

import (
	"flag"
	"fmt"
	"github.com/yasukun/maketree/mktree"
	"os"
)

var input = flag.String("i", "", "input file")
var test = flag.Bool("test", false, "test mode")

// main ...
func main() {
	flag.Parse()

	_, err := os.Stat(*input)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "no such file: %v\n", err)
		os.Exit(1)
	}

	text, err := mktree.ScanAndClean(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unknown error: %v\n", err)
		os.Exit(2)
	}
	mktree.MakeTree(&text, *test)

}
