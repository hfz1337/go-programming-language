// Fetch a URL and print the http response content to stdout.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <url>", os.Args[0])
		os.Exit(0)
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
