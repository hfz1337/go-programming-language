package main

import (
	"fmt"
	"os"
	"strings"
)

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	s = s[strings.LastIndex(s, "/")+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <path>\n", os.Args[0])
		os.Exit(0)
	}
	fmt.Println(basename(os.Args[1]))
}
