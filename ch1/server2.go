// Minimal echo web server that returns the path component of the URL when the / route
// is requested, and returns the total number of requests made to the / route when
// requesting the /count route.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var n int

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/count", count)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	n++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func count(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", n)
	mu.Unlock()
}
