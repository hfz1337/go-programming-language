// Print the index and value of each of the command-line arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i+1, arg)
	}
}
