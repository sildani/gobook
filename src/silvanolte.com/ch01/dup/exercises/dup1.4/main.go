// Modify dup2 to print the names of all files in which
// each duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

type data struct {
	count    int
	filename string
}

func main() {
	counts := make(map[string]data)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.count, line, n.filename)
		}
	}
}

func countLines(f *os.File, counts map[string]data) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		x := counts[input.Text()]
		x.count++
		x.filename += string(f.Name()) + " "
		counts[input.Text()] = x
		// when dealing with command line user input on OS X, use Ctrl-D
		// to signal "end of file", which will cause program to break out
		// of scanning loop
	}
	// note: not accounting for errors stemming from input.Scan() calls
}
