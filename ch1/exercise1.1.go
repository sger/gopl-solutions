package main

import (
	"fmt"
	"os"
	"strings"
)

// Modify the echo program to also print os.Args[0] the name of the command that invoked it.
func main() {
	s, sep := "", ""

	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

	// sort syntax
	fmt.Println(strings.Join(os.Args[0:], " "))
}
