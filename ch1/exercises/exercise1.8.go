// Fetch a URL and print the http response content to stdout.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	var url string
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <url>", os.Args[0])
		os.Exit(0)
	}
	if !strings.HasPrefix(os.Args[1], "https://") && !strings.HasPrefix(os.Args[1], "http://") {
		url = "http://" + os.Args[1]
	} else {
		url = os.Args[1]
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}
