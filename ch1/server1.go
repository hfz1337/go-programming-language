// Minimal echo web server that returns the path component of the URL.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
