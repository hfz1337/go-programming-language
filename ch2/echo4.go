// Print command-line arguments separated an arbitrary separator.
package main

import (
	"flag"
	"fmt"
	"strings"
)

// n and sep are pointers
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
