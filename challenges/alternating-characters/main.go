package main

import (
	"fmt"
	"io"
	"os"
)

func uniqDashC(s string) int {
	var prev rune
	var count int
	for i, c := range s {
		if i == 0 {
			prev = c
			continue
		} else if prev == c {
			count++
		}
		prev = c
	}
	return count
}

func solve(r io.Reader, w io.Writer) {

	var n int
	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(r, &s)
		fmt.Fprintf(w, "%d\n", uniqDashC(s))
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
