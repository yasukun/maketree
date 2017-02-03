package mktree

import (
	"testing"

	"log"
)

// TestCleantree ...
func TestCleantree(t *testing.T) {
	cleantree("    │   ├── 201310_00010_2013-230814.txt.csv")
}

// TestScanAndClean ...
func TestScanAndClean(t *testing.T) {
	text, err := ScanAndClean("../testdata/tree.txt")
	if err != nil {
		t.Error(err)
	}
	log.Println(text)
	MakeTree(&text, false)
}

// TestConutlspasce ...
func TestConutlspasce(t *testing.T) {
	i := countlspace("        workflow")
	if i == 0 {
		t.Error("wrong value")
	}
	log.Println(i)
}
