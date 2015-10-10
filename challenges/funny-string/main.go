package main

import (
	"fmt"
	"io"
	"os"
)

func solve(r io.Reader, w io.Writer) {
	var n int

	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(r, &s)

		ln := len(s)
		funny := true

		for j := 1; j < ln; j++ {
			var x, y byte
			if s[j] > s[j-1] {
				x = s[j] - s[j-1]
			} else {
				x = s[j-1] - s[j]
			}
			if s[ln-j-1] > s[ln-j] {
				y = s[ln-j-1] - s[ln-j]
			} else {
				y = s[ln-j] - s[ln-j-1]
			}
			if x != y {
				funny = false
				break
			}
		}
		if funny {
			fmt.Fprintln(w, "Funny")
		} else {
			fmt.Fprintln(w, "Not Funny")
		}
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
