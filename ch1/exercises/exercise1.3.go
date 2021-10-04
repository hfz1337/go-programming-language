// Evaluate the perfomance of the 3 version of echo.
// Run with: `go run exercice1.3.go argument{1..100}`, for example.
// Example output:
//   echo1 took on average 11417 ns
//   echo2 took on average 11449 ns
//   echo3 took on average 955 ns
// Clearly, using strings.Join is way more faster.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	return strings.Join(os.Args[1:], " ")
}

func main() {
	var start time.Time
	var elapsed int64
	n := 1000000
	// Test echo1
	elapsed = 0
	for i := 1; i < n; i++ {
		start = time.Now()
		_ = echo1()
		elapsed += time.Since(start).Nanoseconds()
	}
	fmt.Printf("echo1 took on average %d ns\n", elapsed/int64(n))

	// Test echo2
	elapsed = 0
	for i := 1; i < n; i++ {
		start = time.Now()
		_ = echo2()
		elapsed += time.Since(start).Nanoseconds()
	}
	fmt.Printf("echo2 took on average %d ns\n", elapsed/int64(n))

	// Test echo3
	elapsed = 0
	for i := 1; i < n; i++ {
		start = time.Now()
		_ = echo3()
		elapsed += time.Since(start).Nanoseconds()
	}
	fmt.Printf("echo3 took on average %d ns\n", elapsed/int64(n))
}
