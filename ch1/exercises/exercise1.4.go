// Print each line that appears more than once in the input, preceded by its count.
// It reads from stdin or a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	occurence := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, occurence)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, occurence)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			if len(files) == 0 {
				fmt.Printf("%d\t%s\n", n, line)
			} else {
				fmt.Printf(
					"%d\t%s (%s)\n", n, line, strings.Join(occurence[line], ", "),
				)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, occurence map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !contains(occurence[input.Text()], f.Name()) {
			occurence[input.Text()] = append(occurence[input.Text()], f.Name())
		}
	}
}

func contains(s []string, target string) bool {
	for _, e := range s {
		if e == target {
			return true
		}
	}
	return false
}
