package main

import (
	"fmt"
	"io"
	"os"
)

func solve(r io.Reader, w io.Writer) {
	var t int

	fmt.Fscan(r, &t)

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		if i > 0 {
			fmt.Fprint(w, "\n")
		}
		fmt.Fprint(w, a+b)
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
