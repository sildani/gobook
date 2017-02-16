package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
    counts[input.Text()]++
    // when dealing with command line user input on OS X, use Ctrl-D
    // to signal "end of file", which will cause program to break out
    // of scanning loop
	}
	// note: not accounting for errors stemming from input.Scan() calls
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
