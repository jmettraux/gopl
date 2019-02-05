package main

import (
	"bufio"
	"fmt"
	"os"
)

func hasDuplicateLines(path string) {
	seenLines := make(map[string]bool)
	f, r := os.Open(path)
	if r != nil {
		fmt.Fprintf(os.Stderr, "%s: couldn't read file %s\n", os.Args[0], path)
	} else {
		s := bufio.NewScanner(f)
		for s.Scan() {
			if len(s.Text()) > 0 && seenLines[s.Text()] {
				fmt.Printf("%s: >%s<\n", path, s.Text())
				break
			} else {
				seenLines[s.Text()] = true
			}
		}
	}
}

func main() {
	for _, path := range os.Args[1:] {
		hasDuplicateLines(path)
	}
}
