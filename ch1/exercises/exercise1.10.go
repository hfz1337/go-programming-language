// Fetch URLs concurrently and report the responses sizes and elapsed times.
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for idx, page_url := range os.Args[1:] {
		// start a goroutine
		go fetch(page_url, ch, idx)
	}
	for range os.Args[1:] {
		// receive from channel
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(page_url string, ch chan<- string, idx int) {
	start := time.Now()
	prased_url, err := url.Parse(page_url)
	resp, err := http.Get(page_url)
	if err != nil {
		// send to channel
		ch <- fmt.Sprint(err)
		return
	}
	f, err := os.Create(fmt.Sprintf("%s_out.txt", prased_url.Hostname()))
	if err != nil {
		ch <- fmt.Sprintf("while creating file: %v", page_url, err)
		return
	}
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()
	f.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", page_url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, page_url)
}
