// Prints execution time from each of solution for echo.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Elapsed times (ms)")
	increaseNumberOfArgs()
	solutionOne()
	solutionTwo()
	solutionThree()
}
func increaseNumberOfArgs() {
	for i := 0; i < 100000; i++ {
		os.Args = append(os.Args, "word")
	}
}
func solutionOne() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("Solution one:", time.Since(start))
}
func solutionTwo() {
	start := time.Now()
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("Solution two:", time.Since(start))
}

func solutionThree() {
	start := time.Now()
	strings.Join(os.Args[1:], " ")
	fmt.Println("Solution three:", time.Since(start))
}
