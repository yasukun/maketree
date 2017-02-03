package mktree

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type FileType int

const (
	Regularfile FileType = iota
	Directory
)

// countlspace ...
func countlspace(line string) int {
	rline := []rune(line)
	count := 0
	for _, r := range rline {
		if r != 32 {
			return count
		}
		count += 1
	}
	return count
}

// filearrjoin ...
func filearrjoin(arr *[]string) string {
	ret := ""
	for _, v := range *arr {
		if v == "" {
			break
		}
		ret = filepath.Join(ret, v)
	}
	return ret
}

// maxlspaces ...
func maxhierarchy(cleantxt *string) int {
	max := 0
	for _, line := range strings.Split(*cleantxt, "\n") {
		max = int(math.Max(float64(max), float64(countlspace(line))))
	}
	return max / 4
}

// fileordirectory ...
func fileordirectory(path string) FileType {
	ext := filepath.Ext(path)
	if len(ext) == 0 {
		return Directory
	} else {
		return Regularfile
	}
}

// touch ...
func touch(path string) error {
	fp, err := os.OpenFile(path, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	return nil
}

// MakeTree ...
func MakeTree(cleantxt *string, test bool) {
	hierarchy := maxhierarchy(cleantxt)
	cols := make([]string, hierarchy+1)
	for _, line := range strings.Split(*cleantxt, "\n") {
		scount := countlspace(line)
		cur := int(scount / 4)
		cols[cur] = strings.Trim(line, " ")
		for i := cur + 1; i < len(cols); i++ {
			cols[i] = ""
		}

		mkfile := filearrjoin(&cols)
		if fileordirectory(mkfile) == Directory {
			fmt.Printf("mkdir -p %s\n", mkfile)
			if !test {
				if err := os.MkdirAll(mkfile, 0744); err != nil {
					fmt.Errorf("%v\n", err)
				}
			}
		} else if fileordirectory(mkfile) == Regularfile {
			fmt.Printf("touch %s\n", mkfile)
			if !test {
				if err := touch(mkfile); err != nil {
					fmt.Errorf("%v\n", err)
				}
			}
		}

	}
}

// cleantree ...
func cleantree(line string) string {
	var flg bool = true
	newrune := []rune{}
	runeline := []rune(line)
	runelen := float64(len(runeline) - 1)
	for pos, c := range runeline {
		nextpos := int32(math.Min(float64(pos+1), runelen))
		if flg {
			newrune = append(newrune, rune(32))
		} else {
			newrune = append(newrune, rune(c))
		}
		if rune(c) == 9472 && runeline[nextpos] == 32 {
			flg = false
		}
	}
	return string(newrune)
}

// ScanAndClean ...
func ScanAndClean(treetxtpath string) (string, error) {
	var buffer bytes.Buffer
	fp, err := os.Open(treetxtpath)
	if err != nil {
		return "", err
	}
	reader := bufio.NewReader(fp)
	linecount := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
		if linecount == 0 {
			buffer.WriteString(string(line) + "\n")
			linecount += 1
			continue
		}
		cleantxt := cleantree(string(line))
		if len(strings.Trim(cleantxt, " ")) == 0 {
			continue
		}
		buffer.WriteString(cleantxt + "\n")

	}
	return strings.TrimRight(buffer.String(), "\n"), nil
}
