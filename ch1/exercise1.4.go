package main

import (
	"bufio"
	"fmt"
	"os"
)

// Modify Dup2 to print the names of all files in which each duplicated line occurs.

func main() {
	counts := make(map[string]int)
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

			for line, n := range counts {
				if n > 1 {
					fmt.Printf("%s\t%d\t%s\n", arg, n, line)
				}
			}
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
